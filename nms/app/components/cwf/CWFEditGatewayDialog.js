/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @flow strict-local
 * @format
 */

import type {
  cwf_gateway,
  magmad_gateway_configs,
} from '../../../generated/MagmaAPIBindings';

import AppBar from '@material-ui/core/AppBar';
import Button from '../../../fbc_js_core/ui/components/design-system/Button';
import CWFGatewayConfigFields from './CWFGatewayConfigFields';
import Dialog from '@material-ui/core/Dialog';
import DialogActions from '@material-ui/core/DialogActions';
import DialogContent from '@material-ui/core/DialogContent';
import GatewayCommandFields from '../GatewayCommandFields';
import MagmaDeviceFields from '../MagmaDeviceFields';
import MagmaV1API from '../../../generated/WebClient';
import React from 'react';
import Tab from '@material-ui/core/Tab';
import Tabs from '@material-ui/core/Tabs';

import nullthrows from '../../../fbc_js_core/util/nullthrows';
import {makeStyles} from '@material-ui/styles';
import {useEnqueueSnackbar} from '../../../fbc_js_core/ui/hooks/useSnackbar';
import {useParams} from 'react-router-dom';
import {useState} from 'react';

const useStyles = makeStyles(_ => ({
  appBar: {
    backgroundColor: '#f5f5f5',
    marginBottom: '20px',
  },
}));

type Props = {
  gateway: cwf_gateway,
  onCancel: () => void,
  onSave: cwf_gateway => void,
};

export default function (props: Props) {
  const [tab, setTab] = useState(0);
  const [magmaConfigs, setMagmaConfigs] = useState(props.gateway.magmad);
  const [allowedGREPeers, setAllowedGREPeers] = useState(
    props.gateway.carrier_wifi.allowed_gre_peers,
  );
  const [ipdrExportDst, setIPDRExportDst] = useState(
    props.gateway.carrier_wifi.ipdr_export_dst,
  );

  const classes = useStyles();
  const params = useParams();
  const enqueueSnackbar = useEnqueueSnackbar();

  const gatewayID = nullthrows(params.gatewayID);
  const networkID = nullthrows(params.networkId);
  const onSave = async () => {
    try {
      await MagmaV1API.putCwfByNetworkIdGatewaysByGatewayId({
        networkId: networkID,
        gatewayId: gatewayID,
        gateway: {
          ...props.gateway,
          carrier_wifi: {
            allowed_gre_peers: allowedGREPeers,
            ipdr_export_dst: ipdrExportDst,
          },
          magmad: getMagmaConfigs(magmaConfigs),
        },
      });
      props.onSave(
        await MagmaV1API.getCwfByNetworkIdGatewaysByGatewayId({
          networkId: networkID,
          gatewayId: gatewayID,
        }),
      );
    } catch (e) {
      enqueueSnackbar(e?.response?.data?.message || e?.message || e, {
        variant: 'error',
      });
    }
  };

  let content;
  switch (tab) {
    case 0:
      content = (
        <CWFGatewayConfigFields
          allowedGREPeers={allowedGREPeers}
          onChange={setAllowedGREPeers}
          ipdrExportDst={ipdrExportDst}
          onIPDRChanged={setIPDRExportDst}
        />
      );
      break;
    case 1:
      content = (
        <MagmaDeviceFields
          configs={magmaConfigs}
          configChangeHandler={(fieldName, value) =>
            setMagmaConfigs({
              ...magmaConfigs,
              [fieldName]: value,
            })
          }
        />
      );
      break;
    case 2:
      content = (
        <GatewayCommandFields
          gatewayID={gatewayID}
          showRestartCommand={true}
          showRebootEnodebCommand={false}
          showPingCommand={true}
          showGenericCommand={true}
        />
      );
      break;

    default:
      // should never happen
      content = <div />;
  }

  return (
    <Dialog open={true} onClose={props.onCancel} maxWidth="md" scroll="body">
      <AppBar position="static" className={classes.appBar}>
        <Tabs
          indicatorColor="primary"
          textColor="primary"
          value={tab}
          onChange={(_, tab) => setTab(tab)}>
          <Tab label="Carrier Wifi" />
          <Tab label="Controller" />
          <Tab label="Command" />
        </Tabs>
      </AppBar>
      <DialogContent>{content}</DialogContent>
      <DialogActions>
        <Button onClick={props.onCancel} skin="regular">
          Cancel
        </Button>
        <Button onClick={onSave}>Save</Button>
      </DialogActions>
    </Dialog>
  );
}

function getMagmaConfigs(
  magmaConfigs: magmad_gateway_configs,
): magmad_gateway_configs {
  return {
    ...magmaConfigs,
    autoupgrade_poll_interval: parseInt(magmaConfigs.autoupgrade_poll_interval),
    checkin_interval: parseInt(magmaConfigs.checkin_interval),
    checkin_timeout: parseInt(magmaConfigs.checkin_timeout),
  };
}
