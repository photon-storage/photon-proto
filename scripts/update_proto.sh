#!/bin/bash

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

# The pre-defined relation map between .proto and git repo, protoc need this relation for
# compile .pb.go and .gw.go
PROTO_MAP=Mgoogle/protobuf/empty.proto=google.golang.org/protobuf/types/known/emptypb,\
Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor,\
Mgoogle/protobuf/timestamp.proto=github.com/golang/protobuf/ptypes/timestamp,\
Mgoogle/api/annotations.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mgoogle/api/client.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mgoogle/api/field_behavior.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mgoogle/api/http.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mgoogle/api/resource.proto=google.golang.org/genproto/googleapis/api/annotations,\
Mproto/gateway/event_source.proto=github.com/grpc-ecosystem/grpc-gateway/v2/proto/gateway,\
Mproto/ext/options.proto=github.com/photon-storage/photon-proto/ext

# the folder for download/compiler the packages like protoc, fast ssz.
PROTO_BUILD="${REPO}/.env/proto_build"
#const version for protoc
PROTOC_VERSION=3.15.8
PROTOC=${PROTO_BUILD}/protoc_${PROTOC_VERSION}_${PLATFORM}/bin/protoc
GOIMPORTS_BIN=${PROTO_BUILD}/gobin/goimports

# download all the required tools and dependencies
if [ "$1" == "install" ]; then
  rm -rf ${PROTO_BUILD} || true
  mkdir -p ${PROTO_BUILD}

  if [ "${PLATFORM}" == "MacOS" ]; then
    # download protobuf bin, will be upgraded to support different environment
    wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip -P ${PROTO_BUILD}
    unzip -q ${PROTO_BUILD}/protoc-${PROTOC_VERSION}-osx-x86_64.zip -d ${PROTO_BUILD}/protoc_${PROTOC_VERSION}_${PLATFORM}
    rm ${PROTO_BUILD}/protoc-${PROTOC_VERSION}-osx-x86_64.zip
  elif [ "${PLATFORM}" == "Linux" ]; then
    wget https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -P ${PROTO_BUILD}
    unzip -q ${PROTO_BUILD}/protoc-${PROTOC_VERSION}-linux-x86_64.zip -d ${PROTO_BUILD}/protoc_${PROTOC_VERSION}_${PLATFORM}
    rm ${PROTO_BUILD}/protoc-${PROTOC_VERSION}-linux-x86_64.zip
  fi

  # download .proto files that are supported by protobuf offical team
  wget https://github.com/protocolbuffers/protobuf/archive/refs/tags/v${PROTOC_VERSION}.zip -P ${PROTO_BUILD}
  unzip -q ${PROTO_BUILD}/v${PROTOC_VERSION}.zip -d ${PROTO_BUILD}
  rm ${PROTO_BUILD}/v${PROTOC_VERSION}.zip

  # download .proto files that are supported by google team
  wget https://github.com/googleapis/googleapis/archive/refs/heads/master.zip -P ${PROTO_BUILD}
  unzip -q ${PROTO_BUILD}/master.zip -d ${PROTO_BUILD}
  rm ${PROTO_BUILD}/master.zip

  # download and compile plugin for generating gw.go
  wget https://github.com/grpc-ecosystem/grpc-gateway/archive/refs/heads/master.zip -P ${PROTO_BUILD}
  unzip -q ${PROTO_BUILD}/master.zip -d ${PROTO_BUILD}
  rm ${PROTO_BUILD}/master.zip
  (cd ${PROTO_BUILD}/grpc-gateway-master/protoc-gen-grpc-gateway && go build -o protoc-gen-grpc-gateway)

# download and compile plugin for generating pb.go
  wget https://github.com/photon-storage/protoc-gen-go-cast/archive/refs/heads/master.zip -P ${PROTO_BUILD}
  unzip -q ${PROTO_BUILD}/master.zip -d ${PROTO_BUILD}
  rm ${PROTO_BUILD}/master.zip
  (cd ${PROTO_BUILD}/protoc-gen-go-cast-master && go build -o protoc-gen-go-cast main.go grpc.go cast.go)

  # download and compile plugin for generating ssz.go
  wget https://github.com/photon-storage/fastssz/archive/refs/heads/master.zip -P ${PROTO_BUILD}
  unzip -q ${PROTO_BUILD}/master.zip -d ${PROTO_BUILD}
  rm ${PROTO_BUILD}/master.zip
  (cd ${PROTO_BUILD}/fastssz-master/sszgen && go build -o sszgen main.go marshal.go tree.go unmarshal.go validate.go size.go tag_parser.go hash.go)

  go install golang.org/x/tools/cmd/goimports@v0.1.9
  mkdir -p $(dirname ${GOIMPORTS_BIN})
  cp $(which goimports) ${GOIMPORTS_BIN}

  exit 0
fi

# generating *.pb.go files
echo "start to generate *.pb.go files"
generate_pb() {
  OUT_DIR="${PROTO_BUILD}/out"
  rm -rf ${OUT_DIR}
  mkdir -p ${OUT_DIR}

  ${PROTOC} --plugin ${PROTO_BUILD}/protoc-gen-go-cast-master/protoc-gen-go-cast \
    --go-cast_out=plugins=grpc,${PROTO_MAP}:${OUT_DIR} ${REPO}/${PKG}/*.proto \
    --proto_path=${REPO} \
    --proto_path=${PROTO_BUILD}/protobuf-${PROTOC_VERSION}/src \
    --proto_path=${PROTO_BUILD}/googleapis-master \
    --proto_path=${PROTO_BUILD}/grpc-gateway-master

  rm -f ${REPO}/${PKG}/*.pb.go
  cp ${OUT_DIR}/github.com/photon-storage/photon-proto/${PKG}/*.pb.go ${REPO}/${PKG}/
  rm -rf ${OUT_DIR}
}

# generating pbs for different package
PKG=ext generate_pb
#PKG=testing generate_pb
PKG=consensus generate_pb
PKG=consensus/validator-client generate_pb
PKG=depot generate_pb
PKG=auditor generate_pb
PKG=keymanager generate_pb

# generating generated.ssz.go files
echo 'start to generate generated.ssz.go files'
generate_ssz() {
  # ssz needs to run in the sandbox or it can crash!
  rm -f ${REPO}/${PKG}/generated.ssz.go

  mkdir -p ${PROTO_BUILD}/ssz_tmp/${PKG}
  cp ${REPO}/${PKG}/*.go ${PROTO_BUILD}/ssz_tmp/${PKG}

  ${PROTO_BUILD}/fastssz-master/sszgen/sszgen --path ${PROTO_BUILD}/ssz_tmp/${PKG} \
  --output ${REPO}/${PKG}/generated.ssz.go --objs ${OBJS}
}

rm -rf ${PROTO_BUILD}/ssz_tmp

# generating ssz for different packages
OBJS=AggregateAttestationAndProof,SignedAggregateAttestationAndProof,SignedBlock,State,SigningData,Status,BlocksByRangeRequest,ENRForkID,Metadata,SignedTransaction,StorageContract,Account,Validator,PoR,Audit
PKG=consensus && generate_ssz

rm -rf ${PROTO_BUILD}/ssz_tmp

${GOIMPORTS_BIN} -w ${REPO}
gofmt -s -w ${REPO}/

echo "start to generate *.pb.gw.go files"
# clean all *.pb.gw.go files
rm -f ${REPO}/consensus/*.gw.go
rm -f ${REPO}/consensus/validator-client/*.gw.go
rm -f ${REPO}/depot/*.gw.go

generate_gw() {
  SRC_FILES=""
  for F in ${GW_FILES} ; do
    echo ">>> ${F}"
    if [ -z "${SRC_FILES}" ]; then
        SRC_FILES="${REPO}/${PKG}/${F}"
    else
        SRC_FILES+=" ${REPO}/${PKG}/${F}"
    fi
  done

  OUT_DIR="${PROTO_BUILD}/out"
  rm -rf ${OUT_DIR}
  mkdir -p ${OUT_DIR}

  ${PROTOC} --plugin ${PROTO_BUILD}/grpc-gateway-master/protoc-gen-grpc-gateway/protoc-gen-grpc-gateway \
  --proto_path=${REPO} \
  --proto_path=${PROTO_BUILD}/protobuf-${PROTOC_VERSION}/src \
  --proto_path=${PROTO_BUILD}/googleapis-master \
  --proto_path=${PROTO_BUILD}/grpc-gateway-master \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true,${PROTO_MAP}:${OUT_DIR} ${SRC_FILES}

  rm -f ${REPO}/${PKG}/*.pb.gw.go
  cp ${OUT_DIR}/github.com/photon-storage/photon-proto/${PKG}/*.pb.gw.go ${REPO}/${PKG}/
  rm -rf ${OUT_DIR}
}

# generating gw for different packages
PKG="consensus" GW_FILES="account.proto info.proto node.proto debug.proto finalized_block_root_container.proto health.proto slasher.proto validator.proto p2p_messages.proto attestation.proto block.proto state.proto transaction.proto account.proto" generate_gw
PKG="consensus/validator-client" GW_FILES="web_api.proto" generate_gw
PKG="depot" GW_FILES="depot.proto" generate_gw

echo -e "\033[0;32mcompiled proto files successfully!\033[0m"
