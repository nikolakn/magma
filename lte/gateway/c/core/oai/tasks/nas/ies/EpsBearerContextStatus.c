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

#include <stdint.h>

#include "lte/gateway/c/core/oai/common/TLVEncoder.h"
#include "lte/gateway/c/core/oai/common/TLVDecoder.h"
#include "lte/gateway/c/core/oai/tasks/nas/ies/EpsBearerContextStatus.h"

//------------------------------------------------------------------------------
int decode_eps_bearer_context_status(
    eps_bearer_context_status_t* epsbearercontextstatus, uint8_t iei,
    uint8_t* buffer, uint32_t len) {
  int decoded = 0;
  uint8_t ielen = 0;

  if (iei > 0) {
    CHECK_IEI_DECODER(iei, *buffer);
    decoded++;
  }

  ielen = *(buffer + decoded);
  decoded++;
  CHECK_LENGTH_DECODER(len - decoded, ielen);
  // IES_DECODE_U16(*epsbearercontextstatus, *(buffer + decoded));
  IES_DECODE_U16(buffer, decoded, *epsbearercontextstatus);
  return decoded;
}

//------------------------------------------------------------------------------
int encode_eps_bearer_context_status(
    eps_bearer_context_status_t* epsbearercontextstatus, uint8_t iei,
    uint8_t* buffer, uint32_t len) {
  uint8_t* lenPtr;
  uint32_t encoded = 0;

  /*
   * Checking IEI and pointer
   */
  CHECK_PDU_POINTER_AND_LENGTH_ENCODER(
      buffer, EPS_BEARER_CONTEXT_STATUS_MINIMUM_LENGTH, len);

  if (iei > 0) {
    *buffer = iei;
    encoded++;
  }

  lenPtr = (buffer + encoded);
  encoded++;
  IES_ENCODE_U16(buffer, encoded, *epsbearercontextstatus);
  *lenPtr = encoded - 1 - ((iei > 0) ? 1 : 0);
  return encoded;
}
