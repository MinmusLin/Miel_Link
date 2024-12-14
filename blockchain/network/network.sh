#!/bin/bash
#
# Copyright IBM Corp All Rights Reserved
#
# SPDX-License-Identifier: Apache-2.0
#

ROOTDIR=$(cd "$(dirname "$0")" && pwd)

export PATH=${PWD}/bin:$PATH
export FABRIC_CFG_PATH=${PWD}/configtx
export VERBOSE=false

pushd ${ROOTDIR} > /dev/null

trap "popd > /dev/null" EXIT

. scripts/utils.sh

: ${CONTAINER_CLI:="docker"}

if command -v ${CONTAINER_CLI}-compose > /dev/null 2>&1; then
    : ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI}-compose"}
else
    : ${CONTAINER_CLI_COMPOSE:="${CONTAINER_CLI} compose"}
fi

infoln "Using ${CONTAINER_CLI} and ${CONTAINER_CLI_COMPOSE}"

function clearContainers() {
    infoln "Removing remaining containers"
    ${CONTAINER_CLI} rm -f $(${CONTAINER_CLI} ps -aq --filter label=service=hyperledger-fabric) 2> /dev/null || true
    ${CONTAINER_CLI} rm -f $(${CONTAINER_CLI} ps -aq --filter name='dev-peer*') 2> /dev/null || true
    ${CONTAINER_CLI} kill "$(${CONTAINER_CLI} ps -q --filter name=ccaas)" 2> /dev/null || true
}

function removeUnwantedImages() {
    infoln "Removing generated chaincode docker images"
    ${CONTAINER_CLI} image rm -f $(${CONTAINER_CLI} images -aq --filter reference='dev-peer*') 2> /dev/null || true
}

NONWORKING_VERSIONS="^1\.0\. ^1\.1\. ^1\.2\. ^1\.3\. ^1\.4\."

function checkPrereqs() {
    peer version > /dev/null 2>&1

    if [[ $? -ne 0 || ! -d "./config" ]]; then
        errorln "Peer binary and configuration files not found.."
        errorln
        errorln "Follow the instructions in the Fabric docs to install the Fabric Binaries:"
        errorln "https://hyperledger-fabric.readthedocs.io/en/latest/install.html"
        exit 1
    fi

    LOCAL_VERSION=$(peer version | sed -ne 's/^ Version: //p')
    DOCKER_IMAGE_VERSION=$(${CONTAINER_CLI} run --rm hyperledger/fabric-tools:latest peer version | sed -ne 's/^ Version: //p')

    infoln "LOCAL_VERSION=$LOCAL_VERSION"
    infoln "DOCKER_IMAGE_VERSION=$DOCKER_IMAGE_VERSION"

    if [ "$LOCAL_VERSION" != "$DOCKER_IMAGE_VERSION" ]; then
        warnln "Local fabric binaries and docker images are out of  sync. This may cause problems."
    fi

    for UNSUPPORTED_VERSION in $NONWORKING_VERSIONS; do
        infoln "$LOCAL_VERSION" | grep -q $UNSUPPORTED_VERSION
        if [ $? -eq 0 ]; then
            fatalln "Local Fabric binary version of $LOCAL_VERSION does not match the versions supported by the test network."
        fi

        infoln "$DOCKER_IMAGE_VERSION" | grep -q $UNSUPPORTED_VERSION
        if [ $? -eq 0 ]; then
            fatalln "Fabric Docker image version of $DOCKER_IMAGE_VERSION does not match the versions supported by the test network."
        fi
    done

    if [ "$CRYPTO" == "cfssl" ]; then
        cfssl version > /dev/null 2>&1
        if [[ $? -ne 0 ]]; then
            errorln "cfssl binary not found.."
            errorln
            errorln "Follow the instructions to install the cfssl and cfssljson binaries:"
            errorln "https://github.com/cloudflare/cfssl#installation"
            exit 1
        fi
    fi

    if [ "$CRYPTO" == "Certificate Authorities" ]; then
        fabric-ca-client version > /dev/null 2>&1
        if [[ $? -ne 0 ]]; then
            errorln "fabric-ca-client binary not found.."
            errorln
            errorln "Follow the instructions in the Fabric docs to install the Fabric Binaries:"
            errorln "https://hyperledger-fabric.readthedocs.io/en/latest/install.html"
            exit 1
        fi
        CA_LOCAL_VERSION=$(fabric-ca-client version | sed -ne 's/ Version: //p')
        CA_DOCKER_IMAGE_VERSION=$(${CONTAINER_CLI} run --rm hyperledger/fabric-ca:latest fabric-ca-client version | sed -ne 's/ Version: //p' | head -1)
        infoln "CA_LOCAL_VERSION=$CA_LOCAL_VERSION"
        infoln "CA_DOCKER_IMAGE_VERSION=$CA_DOCKER_IMAGE_VERSION"
        if [ "$CA_LOCAL_VERSION" != "$CA_DOCKER_IMAGE_VERSION" ]; then
            warnln "Local fabric-ca binaries and docker images are out of sync. This may cause problems."
        fi
    fi
}

function createOrgs() {
    if [ -d "organizations/peerOrganizations" ]; then
        rm -Rf organizations/peerOrganizations && rm -Rf organizations/ordererOrganizations
    fi

    if [ "$CRYPTO" == "cryptogen" ]; then
        which cryptogen
        if [ "$?" -ne 0 ]; then
            fatalln "cryptogen tool not found. exiting"
        fi
        infoln "Generating certificates using cryptogen tool"
        infoln "Creating Org1 Identities"
        set -x
        cryptogen generate --config=./organizations/cryptogen/crypto-config-org1.yaml --output="organizations"
        res=$?
        { set +x; } 2> /dev/null
        if [ $res -ne 0 ]; then
            fatalln "Failed to generate certificates..."
        fi
        infoln "Creating Org2 Identities"
        set -x
        cryptogen generate --config=./organizations/cryptogen/crypto-config-org2.yaml --output="organizations"
        res=$?
        { set +x; } 2> /dev/null
        if [ $res -ne 0 ]; then
            fatalln "Failed to generate certificates..."
        fi
        infoln "Creating Orderer Org Identities"
        set -x
        cryptogen generate --config=./organizations/cryptogen/crypto-config-orderer.yaml --output="organizations"
        res=$?
        { set +x; } 2> /dev/null
        if [ $res -ne 0 ]; then
            fatalln "Failed to generate certificates..."
        fi
    fi

    if [ "$CRYPTO" == "cfssl" ]; then
        . organizations/cfssl/registerEnroll.sh
        peer_cert peer peer0.org1.example.com org1
        peer_cert admin Admin@org1.example.com org1
        infoln "Creating Org2 Identities"
        peer_cert peer peer0.org2.example.com org2
        peer_cert admin Admin@org2.example.com org2
        infoln "Creating Orderer Org Identities"
        orderer_cert orderer orderer.example.com
        orderer_cert admin Admin@example.com
    fi

    if [ "$CRYPTO" == "Certificate Authorities" ]; then
        infoln "Generating certificates using Fabric CA"
        ${CONTAINER_CLI_COMPOSE} -f compose/$COMPOSE_FILE_CA -f compose/$CONTAINER_CLI/${CONTAINER_CLI}-$COMPOSE_FILE_CA up -d 2>&1
        . organizations/fabric-ca/registerEnroll.sh
        while :; do
            if [ ! -f "organizations/fabric-ca/org1/tls-cert.pem" ]; then
                sleep 1
            else
                break
            fi
        done
        infoln "Creating Org1 Identities"
        createOrg1
        infoln "Creating Org2 Identities"
        createOrg2
        infoln "Creating Orderer Org Identities"
        createOrderer
    fi

    infoln "Generating CCP files for Org1 and Org2"
    ./organizations/ccp-generate.sh
}

function networkUp() {
    checkPrereqs
    if [ ! -d "organizations/peerOrganizations" ]; then
        createOrgs
    fi
    COMPOSE_FILES="-f compose/${COMPOSE_FILE_BASE} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_BASE}"
    if [ "${DATABASE}" == "couchdb" ]; then
        COMPOSE_FILES="${COMPOSE_FILES} -f compose/${COMPOSE_FILE_COUCH} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_COUCH}"
    fi
    DOCKER_SOCK="${DOCKER_SOCK}" ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} up -d 2>&1
    $CONTAINER_CLI ps -a
    if [ $? -ne 0 ]; then
        fatalln "Unable to start network"
    fi
}

function createChannel() {
    bringUpNetwork="false"
    local bft_true=$1
    if ! $CONTAINER_CLI info > /dev/null 2>&1; then
        fatalln "$CONTAINER_CLI network is required to be running to create a channel"
    fi
    CONTAINERS=($($CONTAINER_CLI ps | grep hyperledger/ | awk '{print $2}'))
    len=$(echo ${#CONTAINERS[@]})
    if [[ $len -ge 4 ]] && [[ ! -d "organizations/peerOrganizations" ]]; then
        echo "Bringing network down to sync certs with containers"
        networkDown
    fi
    [[ $len -lt 4 ]] || [[ ! -d "organizations/peerOrganizations" ]] && bringUpNetwork="true" || echo "Network Running Already"
    if [ $bringUpNetwork == "true" ]; then
        infoln "Bringing up network"
        networkUp
    fi
    scripts/createChannel.sh $CHANNEL_NAME $CLI_DELAY $MAX_RETRY $VERBOSE $bft_true
}

function deployCC() {
    scripts/deployCC.sh $CHANNEL_NAME $CC_NAME $CC_SRC_PATH $CC_SRC_LANGUAGE $CC_VERSION $CC_SEQUENCE $CC_INIT_FCN $CC_END_POLICY $CC_COLL_CONFIG $CLI_DELAY $MAX_RETRY $VERBOSE
    if [ $? -ne 0 ]; then
        fatalln "Deploying chaincode failed"
    fi
}

function deployCCAAS() {
    scripts/deployCCAAS.sh $CHANNEL_NAME $CC_NAME $CC_SRC_PATH $CCAAS_DOCKER_RUN $CC_VERSION $CC_SEQUENCE $CC_INIT_FCN $CC_END_POLICY $CC_COLL_CONFIG $CLI_DELAY $MAX_RETRY $VERBOSE $CCAAS_DOCKER_RUN
    if [ $? -ne 0 ]; then
        fatalln "Deploying chaincode-as-a-service failed"
    fi
}

function packageChaincode() {
    infoln "Packaging chaincode"
    scripts/packageCC.sh $CC_NAME $CC_SRC_PATH $CC_SRC_LANGUAGE $CC_VERSION true
    if [ $? -ne 0 ]; then
        fatalln "Packaging the chaincode failed"
    fi
}

function listChaincode() {
    export FABRIC_CFG_PATH=${PWD}/config
    . scripts/envVar.sh
    . scripts/ccutils.sh
    setGlobals $ORG
    println
    queryInstalledOnPeer
    println
    listAllCommitted
}

function invokeChaincode() {
    export FABRIC_CFG_PATH=${PWD}/config
    . scripts/envVar.sh
    . scripts/ccutils.sh
    setGlobals $ORG
    chaincodeInvoke $ORG $CHANNEL_NAME $CC_NAME $CC_INVOKE_CONSTRUCTOR
}

function queryChaincode() {
    export FABRIC_CFG_PATH=${PWD}/config
    . scripts/envVar.sh
    . scripts/ccutils.sh
    setGlobals $ORG
    chaincodeQuery $ORG $CHANNEL_NAME $CC_NAME $CC_QUERY_CONSTRUCTOR
}

function networkDown() {
    local temp_compose=$COMPOSE_FILE_BASE
    COMPOSE_FILE_BASE=compose-bft-test-net.yaml
    COMPOSE_BASE_FILES="-f compose/${COMPOSE_FILE_BASE} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_BASE}"
    COMPOSE_COUCH_FILES="-f compose/${COMPOSE_FILE_COUCH} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_COUCH}"
    COMPOSE_CA_FILES="-f compose/${COMPOSE_FILE_CA} -f compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_CA}"
    COMPOSE_FILES="${COMPOSE_BASE_FILES} ${COMPOSE_COUCH_FILES} ${COMPOSE_CA_FILES}"
    COMPOSE_ORG3_BASE_FILES="-f addOrg3/compose/${COMPOSE_FILE_ORG3_BASE} -f addOrg3/compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_ORG3_BASE}"
    COMPOSE_ORG3_COUCH_FILES="-f addOrg3/compose/${COMPOSE_FILE_ORG3_COUCH} -f addOrg3/compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_ORG3_COUCH}"
    COMPOSE_ORG3_CA_FILES="-f addOrg3/compose/${COMPOSE_FILE_ORG3_CA} -f addOrg3/compose/${CONTAINER_CLI}/${CONTAINER_CLI}-${COMPOSE_FILE_ORG3_CA}"
    COMPOSE_ORG3_FILES="${COMPOSE_ORG3_BASE_FILES} ${COMPOSE_ORG3_COUCH_FILES} ${COMPOSE_ORG3_CA_FILES}"

    if [ "${CONTAINER_CLI}" == "docker" ]; then
        DOCKER_SOCK=$DOCKER_SOCK ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} ${COMPOSE_ORG3_FILES} down --volumes --remove-orphans
    elif [ "${CONTAINER_CLI}" == "podman" ]; then
        ${CONTAINER_CLI_COMPOSE} ${COMPOSE_FILES} ${COMPOSE_ORG3_FILES} down --volumes
    else
        fatalln "Container CLI  ${CONTAINER_CLI} not supported"
    fi

    COMPOSE_FILE_BASE=$temp_compose

    if [ "$MODE" != "restart" ]; then
        ${CONTAINER_CLI} volume rm docker_orderer.example.com docker_peer0.org1.example.com docker_peer0.org2.example.com
        clearContainers
        removeUnwantedImages
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf system-genesis-block/*.block organizations/peerOrganizations organizations/ordererOrganizations'
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf organizations/fabric-ca/org1/msp organizations/fabric-ca/org1/tls-cert.pem organizations/fabric-ca/org1/ca-cert.pem organizations/fabric-ca/org1/IssuerPublicKey organizations/fabric-ca/org1/IssuerRevocationPublicKey organizations/fabric-ca/org1/fabric-ca-server.db'
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf organizations/fabric-ca/org2/msp organizations/fabric-ca/org2/tls-cert.pem organizations/fabric-ca/org2/ca-cert.pem organizations/fabric-ca/org2/IssuerPublicKey organizations/fabric-ca/org2/IssuerRevocationPublicKey organizations/fabric-ca/org2/fabric-ca-server.db'
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf organizations/fabric-ca/ordererOrg/msp organizations/fabric-ca/ordererOrg/tls-cert.pem organizations/fabric-ca/ordererOrg/ca-cert.pem organizations/fabric-ca/ordererOrg/IssuerPublicKey organizations/fabric-ca/ordererOrg/IssuerRevocationPublicKey organizations/fabric-ca/ordererOrg/fabric-ca-server.db'
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf addOrg3/fabric-ca/org3/msp addOrg3/fabric-ca/org3/tls-cert.pem addOrg3/fabric-ca/org3/ca-cert.pem addOrg3/fabric-ca/org3/IssuerPublicKey addOrg3/fabric-ca/org3/IssuerRevocationPublicKey addOrg3/fabric-ca/org3/fabric-ca-server.db'
        ${CONTAINER_CLI} run --rm -v "$(pwd):/data" busybox sh -c 'cd /data && rm -rf channel-artifacts log.txt *.tar.gz'
    fi
}

. ./network.config

COMPOSE_FILE_BASE=compose-test-net.yaml
COMPOSE_FILE_COUCH=compose-couch.yaml
COMPOSE_FILE_CA=compose-ca.yaml
COMPOSE_FILE_ORG3_BASE=compose-org3.yaml
COMPOSE_FILE_ORG3_COUCH=compose-couch-org3.yaml
COMPOSE_FILE_ORG3_CA=compose-ca-org3.yaml
SOCK="${DOCKER_HOST:-/var/run/docker.sock}"
DOCKER_SOCK="${SOCK##unix://}"
BFT=0

if [[ $# -lt 1 ]]; then
    printHelp
    exit 0
else
    MODE=$1
    shift
fi

if [ "$MODE" == "cc" ] && [[ $# -lt 1 ]]; then
    printHelp $MODE
    exit 0
fi

if [[ $# -ge 1 ]]; then
    key="$1"
    if [[ "$key" == "createChannel" ]]; then
        export MODE="createChannel"
        shift
    elif [[ "$MODE" == "cc" ]]; then
        if [ "$1" != "-h" ]; then
            export SUBCOMMAND=$key
            shift
        fi
    fi
fi

while [[ $# -ge 1 ]]; do
    key="$1"
    case $key in
        -h)
            printHelp $MODE
            exit 0
            ;;
        -c)
            CHANNEL_NAME="$2"
            shift
            ;;
        -bft)
            BFT=1
            ;;
        -ca)
            CRYPTO="Certificate Authorities"
            ;;
        -cfssl)
            CRYPTO="cfssl"
            ;;
        -r)
            MAX_RETRY="$2"
            shift
            ;;
        -d)
            CLI_DELAY="$2"
            shift
            ;;
        -s)
            DATABASE="$2"
            shift
            ;;
        -ccl)
            CC_SRC_LANGUAGE="$2"
            shift
            ;;
        -ccn)
            CC_NAME="$2"
            shift
            ;;
        -ccv)
            CC_VERSION="$2"
            shift
            ;;
        -ccs)
            CC_SEQUENCE="$2"
            shift
            ;;
        -ccp)
            CC_SRC_PATH="$2"
            shift
            ;;
        -ccep)
            CC_END_POLICY="$2"
            shift
            ;;
        -cccg)
            CC_COLL_CONFIG="$2"
            shift
            ;;
        -cci)
            CC_INIT_FCN="$2"
            shift
            ;;
        -ccaasdocker)
            CCAAS_DOCKER_RUN="$2"
            shift
            ;;
        -verbose)
            VERBOSE=true
            ;;
        -org)
            ORG="$2"
            shift
            ;;
        -i)
            IMAGETAG="$2"
            shift
            ;;
        -cai)
            CA_IMAGETAG="$2"
            shift
            ;;
        -ccic)
            CC_INVOKE_CONSTRUCTOR="$2"
            shift
            ;;
        -ccqc)
            CC_QUERY_CONSTRUCTOR="$2"
            shift
            ;;
        *)
            errorln "Unknown flag: $key"
            printHelp
            exit 1
            ;;
    esac
    shift
done

if [[ $BFT -eq 1 && "$CRYPTO" == "Certificate Authorities" ]]; then
    fatalln "This sample does not yet support the use of consensus type BFT and CA together."
fi

if [ $BFT -eq 1 ]; then
    export FABRIC_CFG_PATH=${PWD}/bft-config
    COMPOSE_FILE_BASE=compose-bft-test-net.yaml
fi

if [ ! -d "organizations/peerOrganizations" ]; then
    CRYPTO_MODE="with crypto from '${CRYPTO}'"
else
    CRYPTO_MODE=""
fi

if [ "$MODE" == "prereq" ]; then
    infoln "Installing binaries and fabric images. Fabric Version: ${IMAGETAG}  Fabric CA Version: ${CA_IMAGETAG}"
    installPrereqs
elif [ "$MODE" == "up" ]; then
    infoln "Starting nodes with CLI timeout of '${MAX_RETRY}' tries and CLI delay of '${CLI_DELAY}' seconds and using database '${DATABASE}' ${CRYPTO_MODE}"
    networkUp
elif [ "$MODE" == "createChannel" ]; then
    infoln "Creating channel '${CHANNEL_NAME}'."
    infoln "If network is not up, starting nodes with CLI timeout of '${MAX_RETRY}' tries and CLI delay of '${CLI_DELAY}' seconds and using database '${DATABASE} ${CRYPTO_MODE}"
    createChannel $BFT
elif [ "$MODE" == "down" ]; then
    infoln "Stopping network"
    networkDown
elif [ "$MODE" == "restart" ]; then
    infoln "Restarting network"
    networkDown
    networkUp
elif [ "$MODE" == "deployCC" ]; then
    infoln "deploying chaincode on channel '${CHANNEL_NAME}'"
    deployCC
elif [ "$MODE" == "deployCCAAS" ]; then
    infoln "deploying chaincode-as-a-service on channel '${CHANNEL_NAME}'"
    deployCCAAS
elif [ "$MODE" == "cc" ] && [ "$SUBCOMMAND" == "package" ]; then
    packageChaincode
elif [ "$MODE" == "cc" ] && [ "$SUBCOMMAND" == "list" ]; then
    listChaincode
elif [ "$MODE" == "cc" ] && [ "$SUBCOMMAND" == "invoke" ]; then
    invokeChaincode
elif [ "$MODE" == "cc" ] && [ "$SUBCOMMAND" == "query" ]; then
    queryChaincode
else
    printHelp
    exit 1
fi
