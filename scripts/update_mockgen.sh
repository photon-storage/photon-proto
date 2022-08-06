#!/bin/bash

# Script to update mock files after consensus/*.proto changes.
# Use a space to separate mock destination from its interfaces.

set -e
set -x

case "$(uname -s)" in
  Darwin)
    PLATFORM="MacOS"
    REPO=$(cd "$(dirname "$0")/.."; pwd -P)
    ;;

  Linux)
    PLATFORM="Linux"
    SCRIPT_REALPATH=$(realpath -e -- "${BASH_SOURCE[0]}")
    REPO="${SCRIPT_REALPATH%/*}/.."
    ;;

  CYGWIN*|MINGW32*|MSYS*|MINGW*)
    PLATFORM="Windows"
    echo "Windows NOT supported"
    exit 1
    ;;

  *)
    echo "unknown OS"
    exit 1
    ;;
esac

PROTO_BUILD="${REPO}/.env/proto_build"
MOCKGEN_BIN=${PROTO_BUILD}/gobin/mockgen
GOIMPORTS_BIN=${PROTO_BUILD}/gobin/goimports

if [ "$1" == "install" ]; then
  go install github.com/golang/mock/mockgen@v1.6.0
  mkdir -p $(dirname ${MOCKGEN_BIN})
  cp $(which mockgen) ${MOCKGEN_BIN}

  exit 0
fi

MOCK_PATH="${REPO}/mockgen"
MOCKS=(
      "consensus ${MOCK_PATH}/node_server_mock.go Node_StreamChainHeadServer,Node_StreamAttestationsServer,Node_StreamBlocksServer,Node_StreamValidatorsInfoServer,Node_StreamIndexedAttestationsServer,Node_WaitForActivationServer,Node_WaitForChainStartServer,Node_StreamDutiesServer,Node_StreamFinalizedEpochServer"
      "consensus ${MOCK_PATH}/node_client_mock.go NodeClient,Node_StreamChainHeadClient,Node_StreamAttestationsClient,Node_StreamBlocksClient,Node_StreamValidatorsInfoClient,Node_StreamIndexedAttestationsClient,Node_WaitForChainStartClient,Node_WaitForActivationClient,Node_StreamDutiesClient,Node_StreamFinalizedEpochClient"
      "consensus ${MOCK_PATH}/slasher_client_mock.go SlasherClient"
      "consensus ${MOCK_PATH}/info_service_mock.go InfoClient"
      "depot ${MOCK_PATH}/depot_client_mock.go DepotClient"
      "keymanager ${MOCK_PATH}/keymanager_mock.go KeyManagerClient"
)

for ((i = 0; i < ${#MOCKS[@]}; i++)); do
    PKG=${MOCKS[i]%% *};
    PRE=${MOCKS[i]% *}
    FILE=${PRE#* };
    INTERFACES=${MOCKS[i]##* };
    echo "generating ${FILE} for interfaces: ${INTERFACES}";
    GO11MODULE=on ${MOCKGEN_BIN} -package=mockgen -destination="${FILE}" github.com/photon-storage/photon-proto/${PKG} "${INTERFACES}"
done

${GOIMPORTS_BIN} -w "${MOCK_PATH}/."
gofmt -s -w "${MOCK_PATH}/."

echo -e "\033[0;32mgenerated service mock files successfully!\033[0m"
