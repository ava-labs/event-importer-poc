// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {EVMLog, EVMReceipt, EVMBlockHeader} from "./IEventImporter.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RlpReader.sol";

library RLPUtils {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    function decodeReceipt(bytes memory encodedReceipt) internal pure returns (EVMReceipt memory) {
        RLPReader.RLPItem[] memory receipt = encodedReceipt.toRlpItem().toList();
        require(receipt.length == 4, "Invalid number of RLP elements in receipt");
        EVMReceipt memory evmReceipt;
        evmReceipt.status = receipt[0].toUint() == 1;
        evmReceipt.cululativeGasUsed = uint64(receipt[1].toUint());
        evmReceipt.bloom = receipt[2].toBytes();
        RLPReader.RLPItem[] memory logs = receipt[3].toList();
        evmReceipt.logs = new EVMLog[](logs.length);
        for (uint256 i = 0; i < logs.length; i++) {
            RLPReader.RLPItem[] memory log = logs[i].toList();
            require(log.length == 3, "Invalid number of RLP elements in log");
            evmReceipt.logs[i].loggerAddress = log[0].toAddress();
            RLPReader.RLPItem[] memory topics = log[1].toList();
            evmReceipt.logs[i].topics = new bytes32[](topics.length);
            for (uint256 j = 0; j < topics.length; j++) {
                evmReceipt.logs[i].topics[j] = bytes32(topics[j].toBytes());
            }
            evmReceipt.logs[i].data = log[2].toBytes();
        }
        return evmReceipt;
    }

    function decodeBlockNumberAndReceiptsRoot(bytes memory encodedBlockHeader)
        internal
        pure
        returns (uint256, bytes32)
    {
        // RLP decode the block header.
        RLPReader.RLPItem[] memory blockHeader = encodedBlockHeader.toRlpItem().toList();
        require(blockHeader.length >= 15, "Invalid number of RLP elements in block header");

        // Extract the block number and the receipts root from the RLP encoding.
        uint256 blockNumber = blockHeader[8].toUint();
        bytes32 receiptsRoot = bytes32(blockHeader[5].toBytes());

        return (blockNumber, receiptsRoot);
    }

    function encodeUint256(uint256 index) internal pure returns (bytes memory) {
        if (index == 0) {
            return hex"80";
        }
        // TODO: Support RLP encoding of larger tx indexes.
        require(index <= 127, "Invalid tx index");
        return abi.encodePacked(uint8(index));
    }
}
