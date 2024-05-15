// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {EVMEventInfo, EVMLog, EVMReceipt, EVMBlockHeader, IEventImporter} from "./IEventImporter.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPReader} from "@solidity-merkle-trees/trie/ethereum/RlpReader.sol";

contract EventImporter is IEventImporter {
    using RLPReader for bytes;
    using RLPReader for RLPReader.RLPItem;

    IWarpMessenger public warpMessenger;

    constructor() {
        warpMessenger = IWarpMessenger(0x0200000000000000000000000000000000000005);
    }

    /*
     * @notice 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     */
    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes[] calldata receiptProof, uint256 logIndex)
        external
        view
        returns (EVMEventInfo memory)
    {
        // Get the verified block has via the Warp precompile.
        (WarpBlockHash memory warpBlockHash, bool valid) = warpMessenger.getVerifiedWarpBlockHash(0);
        require(valid, "Invalid WarpBlockHash");

        // Check that the blockHeader matches the authenticated block hash.
        require(keccak256(blockHeader) == warpBlockHash.blockHash, "Invalid blockHeader");

        // RLP decode the required values from the block header.
        (uint256 blockNumber, bytes32 receiptsRoot) = decodeBlockHeader(blockHeader);

        // Construct the key of the trie receipt proof.
        bytes[] memory receiptKeys = new bytes[](1);
        receiptKeys[0] = rlpEncodeTxIndex(txIndex);

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(receiptsRoot, receiptProof, receiptKeys);
        require(results.length == 1, "Invalid number of results in receipt proof");
        require(results[0].value.length > 0, "Invalid receipt proof");

        EVMReceipt memory receipt = decodeReceipt(results[0].value);
        require(logIndex < receipt.logs.length, "Invalid log index");

        return EVMEventInfo({
            blockchainID: warpBlockHash.sourceChainID,
            blockNumber: blockNumber,
            txIndex: txIndex,
            logIndex: logIndex,
            log: receipt.logs[logIndex]
        });
    }

    function decodeReceipt(bytes memory encodedReceipt) public pure returns (EVMReceipt memory) {
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

    function decodeBlockHeader(bytes memory encodedBlockHeader) public pure returns (uint256, bytes32) {
        // RLP decode the block header.
        RLPReader.RLPItem[] memory blockHeader = encodedBlockHeader.toRlpItem().toList();
        require(blockHeader.length >= 15, "Invalid number of RLP elements in block header");

        // Extract the block number and the receipts root from the RLP encoding.
        uint256 blockNumber = blockHeader[8].toUint();
        bytes32 receiptsRoot = bytes32(blockHeader[5].toBytes());

        return (blockNumber, receiptsRoot);
    }

    function rlpEncodeTxIndex(uint256 index) public pure returns (bytes memory) {
        if (index == 0) {
            return hex"80";
        }
        // TODO: Support RLP encoding of larger tx indexes.
        require(index <= 127, "Invalid tx index");
        return abi.encodePacked(uint8(index));
    }
}
