/*
 * Licensed to the OpenAirInterface (OAI) Software Alliance under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The OpenAirInterface Software Alliance licenses this file to You under
 * the terms found in the LICENSE file in the root of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *-------------------------------------------------------------------------------
 * For more information about the OpenAirInterface (OAI) Software Alliance:
 *      contact@openairinterface.org
 */
#include <unistd.h>
#include <stdlib.h>

#include "lte/gateway/c/core/common/common_defs.h"
#include "lte/gateway/c/core/oai/common/log.h"
#include "lte/gateway/c/core/oai/include/pgw_config.h"
#include "lte/gateway/c/core/oai/include/sgw_config.h"
#include "lte/gateway/c/core/oai/lib/bstr/bstrlib.h"
#include "lte/gateway/c/core/oai/tasks/mme_app/mme_app_embedded_spgw.h"

char* USAGE_TEXT =
    "==== EURECOM %s version: %s ====\n"  // PACKAGE_NAME, PACKAGE_VERSION
    "Please report any bug to: %s\n"      // PACKAGE_BUGREPORT
    "Usage: %s [options]\n"               // exe_path
    "Available options:\n"
    "-h      Print this help and return\n"
    "-c <path>\n"
    "        Set the configuration file for mme\n"
    "        See template in UTILS/CONF\n"
    "-s <path>\n"
    "        Set the configuration file for S/P-GW\n"
    "        See template in ETC\n"
    "-K <file>\n"
    "        Output intertask messages to provided file\n"
    "-V      Print %s version and return\n"  // PACKAGE_NAME
    "-v[1-2] Debug level:\n"
    "        1 -> ASN1 XER printf on and ASN1 debug off\n"
    "        2 -> ASN1 XER printf on and ASN1 debug on\n";

static void usage(char* exe_path) {
  OAILOG_INFO(LOG_CONFIG, USAGE_TEXT, PACKAGE_NAME, PACKAGE_VERSION,
              PACKAGE_BUGREPORT, exe_path, PACKAGE_NAME);
}

status_code_e mme_config_embedded_spgw_parse_opt_line(
    int argc, char* argv[], mme_config_t* mme_config_p,
    amf_config_t* amf_config_p, spgw_config_t* spgw_config_p) {
  int c;

  mme_config_init(mme_config_p);
  spgw_config_init(spgw_config_p);
  amf_config_init(amf_config_p);

  while ((c = getopt(argc, argv, "c:hi:Ks:v:V:p:")) != -1) {
    switch (c) {
      case 'c':
        mme_config_p->config_file = bfromcstr(optarg);
        OAILOG_DEBUG(LOG_CONFIG, "mme_config.config_file %s",
                     bdata(mme_config_p->config_file));
        break;

#if MME_BENCHMARK
      case 'p': {
        mme_config_p->test_param = atoi(optarg);
        mme_config_p->test_type = TEST_SERIALIZATION_PROTOBUF;
        mme_config_p->run_mode = RUN_MODE_TEST;
        OAI_FPRINTF_INFO("Test serialization protobuf, parameter %u\n",
                         mme_config_p->test_param);
      } break;
#endif

      case 'v':
        mme_config_p->log_config.asn1_verbosity_level = atoi(optarg);
        break;

      case 'V':
        OAILOG_DEBUG(LOG_CONFIG,
                     "==== EURECOM %s v%s ===="
                     "Please report any bug to: %s",
                     PACKAGE_NAME, PACKAGE_VERSION, PACKAGE_BUGREPORT);

        break;

      case 'K':
        mme_config_p->itti_config.log_file = bfromcstr(optarg);
        spgw_config_p->sgw_config.itti_config.log_file = bfromcstr(optarg);

        OAILOG_DEBUG(LOG_CONFIG, "mme_config.itti_config.log_file %s",
                     bdata(mme_config_p->itti_config.log_file));
        OAILOG_DEBUG(LOG_CONFIG,
                     "spgw_config.sgw_config.itti_config.log_file %s",
                     bdata(spgw_config_p->sgw_config.itti_config.log_file));

        break;

      case 's':
        spgw_config_p->config_file = bfromcstr(optarg);
        spgw_config_p->pgw_config.config_file = bfromcstr(optarg);
        spgw_config_p->sgw_config.config_file = bfromcstr(optarg);

        OAILOG_DEBUG(LOG_CONFIG, "spgw_config.config_file %s\n",
                     bdata(spgw_config_p->config_file));

        break;

      case 'h':
#if !MME_BENCHMARK
      case 'p':
#endif
      default:
        usage(argv[0]);
        exit(0);
        break;
    }
  }

  if (!mme_config_p->config_file) {
    mme_config_p->config_file = bfromcstr("/usr/local/etc/oai/mme.conf");
  }

  if (!amf_config_p->config_file) {
    amf_config_p->config_file = mme_config_p->config_file;
  }

  if (!spgw_config_p->config_file) {
    spgw_config_p->config_file = bfromcstr("/usr/local/etc/oai/spgw.conf");
    spgw_config_p->pgw_config.config_file =
        bfromcstr("/usr/local/etc/oai/spgw.conf");
    spgw_config_p->sgw_config.config_file =
        bfromcstr("/usr/local/etc/oai/spgw.conf");
  }

  if (mme_config_parse_file(mme_config_p) != 0) {
    return RETURNerror;
  }

  if (spgw_config_parse_file(spgw_config_p) != 0) {
    return RETURNerror;
  }

  if (amf_config_parse_file(amf_config_p, mme_config_p) != 0) {
    return RETURNerror;
  }

  amf_config_display(amf_config_p);
  mme_config_display(mme_config_p);
  spgw_config_display(spgw_config_p);

  return RETURNok;
}
