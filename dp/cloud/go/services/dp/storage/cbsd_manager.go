/*
Copyright 2022 The Magma Authors.

This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storage

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"

	"magma/dp/cloud/go/services/dp/storage/db"
	"magma/orc8r/cloud/go/sqorc"
	"magma/orc8r/lib/go/merrors"
)

type CbsdFilter struct {
	SerialNumber string
}

type CbsdManager interface {
	CreateCbsd(networkId string, data *DBCbsd) error
	UpdateCbsd(networkId string, id int64, data *DBCbsd) error
	DeleteCbsd(networkId string, id int64) error
	FetchCbsd(networkId string, id int64) (*DetailedCbsd, error)
	ListCbsd(networkId string, pagination *Pagination, filter *CbsdFilter) (*DetailedCbsdList, error)
}

type DetailedCbsdList struct {
	Cbsds []*DetailedCbsd
	Count int64
}

type DetailedCbsd struct {
	Cbsd       *DBCbsd
	CbsdState  *DBCbsdState
	Grant      *DBGrant
	GrantState *DBGrantState
}

func NewCbsdManager(db *sql.DB, builder sqorc.StatementBuilder, errorChecker sqorc.ErrorChecker) *cbsdManager {
	return &cbsdManager{
		db:           db,
		builder:      builder,
		cache:        &enumCache{cache: map[string]map[string]int64{}},
		errorChecker: errorChecker,
	}
}

type cbsdManager struct {
	db           *sql.DB
	builder      sqorc.StatementBuilder
	cache        *enumCache
	errorChecker sqorc.ErrorChecker
}

type enumCache struct {
	cache map[string]map[string]int64
}

func (c *cbsdManager) CreateCbsd(networkId string, data *DBCbsd) error {
	_, err := sqorc.ExecInTx(c.db, nil, nil, func(tx *sql.Tx) (interface{}, error) {
		runner := c.getInTransactionManager(tx)
		err := runner.createCbsdWithActiveModeConfig(networkId, data)
		return nil, err
	})
	return makeError(err, c.errorChecker)
}

func (c *cbsdManager) UpdateCbsd(networkId string, id int64, data *DBCbsd) error {
	_, err := sqorc.ExecInTx(c.db, nil, nil, func(tx *sql.Tx) (interface{}, error) {
		runner := c.getInTransactionManager(tx)
		err := runner.updateCbsd(networkId, id, data)
		return nil, err
	})
	return makeError(err, c.errorChecker)
}

func (c *cbsdManager) DeleteCbsd(networkId string, id int64) error {
	_, err := sqorc.ExecInTx(c.db, nil, nil, func(tx *sql.Tx) (interface{}, error) {
		runner := c.getInTransactionManager(tx)
		err := runner.markCbsdAsDeleted(networkId, id)
		return nil, err
	})
	return makeError(err, c.errorChecker)
}

func (c *cbsdManager) FetchCbsd(networkId string, id int64) (*DetailedCbsd, error) {
	cbsd, err := sqorc.ExecInTx(c.db, nil, nil, func(tx *sql.Tx) (interface{}, error) {
		runner := c.getInTransactionManager(tx)
		return runner.fetchDetailedCbsd(networkId, id)
	})
	if err != nil {
		return nil, makeError(err, c.errorChecker)
	}
	return cbsd.(*DetailedCbsd), nil
}

func (c *cbsdManager) ListCbsd(networkId string, pagination *Pagination, filter *CbsdFilter) (*DetailedCbsdList, error) {
	cbsds, err := sqorc.ExecInTx(c.db, nil, nil, func(tx *sql.Tx) (interface{}, error) {
		runner := c.getInTransactionManager(tx)
		return runner.listDetailedCbsd(networkId, pagination, filter)
	})
	if err != nil {
		return nil, makeError(err, c.errorChecker)
	}
	return cbsds.(*DetailedCbsdList), nil
}

func (c *cbsdManager) getInTransactionManager(tx sq.BaseRunner) *cbsdManagerInTransaction {
	return &cbsdManagerInTransaction{
		builder: c.builder.RunWith(tx),
		cache:   c.cache,
	}
}

type cbsdManagerInTransaction struct {
	builder sq.StatementBuilderType
	cache   *enumCache
}

func (c *cbsdManagerInTransaction) createCbsdWithActiveModeConfig(networkId string, data *DBCbsd) error {
	unregisteredState, err := c.cache.getValue(c.builder, &DBCbsdState{}, "unregistered")
	if err != nil {
		return err
	}
	registeredState, err := c.cache.getValue(c.builder, &DBCbsdState{}, "registered")
	if err != nil {
		return err
	}
	data.StateId = db.MakeInt(unregisteredState)
	data.DesiredStateId = db.MakeInt(registeredState)
	data.NetworkId = db.MakeString(networkId)
	columns := append(getCbsdWriteFields(), "state_id", "desired_state_id", "network_id")
	_, err = db.NewQuery().
		WithBuilder(c.builder).
		From(data).
		Select(db.NewIncludeMask(columns...)).
		Insert()
	return err
}

func (e *enumCache) getValue(builder sq.StatementBuilderType, model db.Model, name string) (int64, error) {
	meta := model.GetMetadata()
	_, ok := e.cache[meta.Table]
	if !ok {
		e.cache[meta.Table] = map[string]int64{}
	}
	if value, ok := e.cache[meta.Table][name]; ok {
		return value, nil
	}
	r, err := db.NewQuery().
		WithBuilder(builder).
		From(model).
		Select(db.NewIncludeMask("id")).
		Where(sq.Eq{"name": name}).
		Fetch()
	if err != nil {
		return 0, err
	}
	e.cache[meta.Table][name] = r[0].(EnumModel).GetId()
	return e.cache[meta.Table][name], nil
}

func getCbsdWriteFields() []string {
	return []string{
		"fcc_id", "cbsd_serial_number", "user_id",
		"min_power", "max_power", "antenna_gain", "number_of_ports",
		"preferred_bandwidth_mhz", "preferred_frequencies_mhz",
	}
}

func (c *cbsdManagerInTransaction) updateCbsd(networkId string, id int64, data *DBCbsd) error {
	if err := c.checkIfCbsdExists(networkId, id); err != nil {
		return err
	}
	data.ShouldDeregister = db.MakeBool(true)
	columns := append(getCbsdWriteFields(), "should_deregister")
	return db.NewQuery().
		WithBuilder(c.builder).
		From(data).
		Select(db.NewIncludeMask(columns...)).
		Where(sq.Eq{"id": id}).
		Update()
}

func (c *cbsdManagerInTransaction) checkIfCbsdExists(networkId string, id int64) error {
	_, err := db.NewQuery().
		WithBuilder(c.builder).
		From(&DBCbsd{}).
		Select(db.NewIncludeMask("id")).
		Where(getCbsdFiltersWithId(networkId, id)).
		Fetch()
	return err
}

func (c *cbsdManagerInTransaction) markCbsdAsDeleted(networkId string, id int64) error {
	if err := c.checkIfCbsdExists(networkId, id); err != nil {
		return err
	}
	return db.NewQuery().
		WithBuilder(c.builder).
		From(&DBCbsd{IsDeleted: db.MakeBool(true)}).
		Select(db.NewIncludeMask("is_deleted")).
		Where(sq.Eq{"id": id}).
		Update()
}

func (c *cbsdManagerInTransaction) fetchDetailedCbsd(networkId string, id int64) (*DetailedCbsd, error) {
	res, err := buildDetailedCbsdQuery(c.builder).
		Where(getCbsdFiltersWithId(networkId, id)).
		Fetch()
	if err != nil {
		return nil, err
	}
	return convertToDetails(res), nil
}

func convertToDetails(models []db.Model) *DetailedCbsd {
	return &DetailedCbsd{
		Cbsd:       models[0].(*DBCbsd),
		CbsdState:  models[1].(*DBCbsdState),
		Grant:      models[2].(*DBGrant),
		GrantState: models[3].(*DBGrantState),
	}
}

func buildDetailedCbsdQuery(builder sq.StatementBuilderType) *db.Query {
	return db.NewQuery().
		WithBuilder(builder).
		From(&DBCbsd{}).
		Select(db.NewExcludeMask("network_id", "state_id", "desired_state_id",
			"is_deleted", "should_deregister", "grant_attempts")).
		Join(db.NewQuery().
			From(&DBCbsdState{}).
			On(db.On(CbsdTable, "state_id", CbsdStateTable, "id")).
			Select(db.NewIncludeMask("name"))).
		Join(db.NewQuery().
			From(&DBGrant{}).
			On(db.On(CbsdTable, "id", GrantTable, "cbsd_id")).
			Select(db.NewIncludeMask(
				"grant_expire_time", "transmit_expire_time",
				"low_frequency", "high_frequency", "max_eirp")).
			Join(db.NewQuery().
				From(&DBGrantState{}).
				On(sq.And{
					db.On(GrantTable, "state_id", GrantStateTable, "id"),
					sq.NotEq{GrantStateTable + ".name": "idle"},
				}).
				Select(db.NewIncludeMask("name"))).
			Nullable())
}

func (c *cbsdManagerInTransaction) listDetailedCbsd(networkId string, pagination *Pagination, filter *CbsdFilter) (*DetailedCbsdList, error) {
	count, err := countCbsds(networkId, c.builder)
	if err != nil {
		return nil, err
	}
	query := buildDetailedCbsdQuery(c.builder)
	res, err := buildPagination(query, pagination).
		Where(getCbsdFilters(networkId, filter)).
		OrderBy(CbsdTable+".id", db.OrderAsc).
		List()
	if err != nil {
		return nil, err
	}
	cbsds := make([]*DetailedCbsd, len(res))
	for i, models := range res {
		cbsds[i] = convertToDetails(models)
	}
	return &DetailedCbsdList{
		Cbsds: cbsds,
		Count: count,
	}, nil
}

func countCbsds(networkId string, builder sq.StatementBuilderType) (int64, error) {
	return db.NewQuery().
		WithBuilder(builder).
		From(&DBCbsd{}).
		Where(getCbsdFilters(networkId, nil)).
		Count()
}

func makeError(err error, checker sqorc.ErrorChecker) error {
	if err == sql.ErrNoRows {
		return merrors.ErrNotFound
	}
	return checker.GetError(err)
}

func getCbsdFiltersWithId(networkId string, id int64) sq.Eq {
	filters := getCbsdFilters(networkId, nil)
	filters[CbsdTable+".id"] = id
	return filters
}

func getCbsdFilters(networkId string, filter *CbsdFilter) sq.Eq {
	filters := sq.Eq{
		CbsdTable + ".network_id": networkId,
		CbsdTable + ".is_deleted": false,
	}
	if filter != nil {
		if filter.SerialNumber != "" {
			filters[CbsdTable+".cbsd_serial_number"] = filter.SerialNumber
		}
	}
	return filters
}
