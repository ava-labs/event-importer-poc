#!/usr/bin/env bash
# Copyright (C) 2024, Ava Labs, Inc. All rights reserved.
# See the file LICENSE for licensing terms.

set -e

REPO_PATH=$(
  cd "$(dirname "${BASH_SOURCE[0]}")"
  cd .. && pwd
)

source $REPO_PATH/scripts/constants.sh
source $REPO_PATH/scripts/versions.sh
source $TELEPORTER_PATH/scripts/utils.sh

setARCH

# Contract names to generate Go bindings for
DEFAULT_CONTRACT_LIST="MockPriceFeedAggregator IEventImporter EventImporter"

CONTRACT_LIST=
HELP=
while [ $# -gt 0 ]; do
    case "$1" in
        -c | --contract) CONTRACT_LIST=$2 ;;
        -h | --help) HELP=true ;;
    esac
    shift
done

if [ "$HELP" = true ]; then
    echo "Usage: ./scripts/abi_bindings.sh [OPTIONS]"
    echo "Build contracts and generate Go bindings"
    echo ""
    echo "Options:"
    echo "  -c, --contract <contract_name>          Generate Go bindings for the contract. If empty, generate Go bindings for a default list of contracts"
    echo "  -c, --contract "contract1 contract2"    Generate Go bindings for multiple contracts"
    echo "  -h, --help                              Print this help message"
    exit 0
fi

if ! command -v forge &> /dev/null; then
    echo "forge not found. You can install by calling $REPO_PATH/scripts/install_foundry.sh" && exit 1
fi

echo "Building subnet-evm abigen"
go install github.com/ava-labs/subnet-evm/cmd/abigen@${SUBNET_EVM_VERSION}

# Force recompile of all contracts to prevent against using previous
# compilations that did not generate new ABI files.
echo "Building Contracts"
cd $REPO_PATH/contracts
forge build --skip test --force --extra-output-files abi bin

contract_names=($CONTRACT_LIST)

# If CONTRACT_LIST is empty, use DEFAULT_CONTRACT_LIST
if [[ -z "${CONTRACT_LIST}" ]]; then
    contract_names=($DEFAULT_CONTRACT_LIST)
fi

cd $REPO_PATH/contracts/src
for contract_name in "${contract_names[@]}"
do
    path=$(find . -name $contract_name.sol)
    dir=$(dirname $path)
    abi_file=$REPO_PATH/contracts/out/$contract_name.sol/$contract_name.abi.json
    if ! [ -f $abi_file ]; then
        echo "Error: Contract $contract_name abi file not found"
        exit 1
    fi

    echo "Generating Go bindings for $contract_name..."
    gen_path=$REPO_PATH/abi-bindings/go/$dir/$contract_name
    mkdir -p $gen_path
    $GOPATH/bin/abigen --abi $abi_file \
                       --pkg $(convertToLower $contract_name) \
                       --bin $REPO_PATH/contracts/out/$contract_name.sol/$contract_name.bin \
                       --type $contract_name \
                       --out $gen_path/$contract_name.go
    echo "Done generating Go bindings for $contract_name."
done

exit 0
