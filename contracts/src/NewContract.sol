// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {EVMEventInfo, EVMLog, EVMReceipt, EVMBlockHeader, IEventImporter} from "./IEventImporter.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {MerklePatricia, StorageValue} from "@solidity-merkle-trees/MerklePatricia.sol";
import {RLPUtils} from "./RLPUtils.sol";

contract NewContract is IEventImporter {
    event WeGotIt();

    IWarpMessenger public warpMessenger;

    constructor() {
        warpMessenger = IWarpMessenger(0x0200000000000000000000000000000000000005);
    }

    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes[] calldata receiptProof, uint256 logIndex)
        external
    {
        // Get the verified block has via the Warp precompile.
        (WarpBlockHash memory warpBlockHash, bool valid) = warpMessenger.getVerifiedWarpBlockHash(0);
        require(valid, "Invalid WarpBlockHash");

        // Check that the blockHeader matches the authenticated block hash.
        require(keccak256(blockHeader) == warpBlockHash.blockHash, "Invalid blockHeader");

        // RLP decode the required values from the block header.
        (uint256 blockNumber, bytes32 receiptsRoot) = RLPUtils.decodeBlockNumberAndReceiptsRoot(blockHeader);

        // Construct the key of the trie receipt proof.
        bytes[] memory receiptKeys = new bytes[](1);
        receiptKeys[0] = RLPUtils.encodeUint256(txIndex);

        // Verify the trie proof against the receipts root.
        StorageValue[] memory results = MerklePatricia.VerifyEthereumProof(receiptsRoot, receiptProof, receiptKeys);
        // require(results.length == 1, "Invalid number of results in receipt proof");
        // require(results[0].value.length > 0, "Invalid receipt proof");

        // EVMReceipt memory receipt = RLPUtils.decodeReceipt(results[0].value);
        // require(logIndex < receipt.logs.length, "Invalid log index");

        // _onEventImport(
        //     EVMEventInfo({
        //         blockchainID: warpBlockHash.sourceChainID,
        //         blockNumber: blockNumber,
        //         txIndex: txIndex,
        //         logIndex: logIndex,
        //         log: receipt.logs[logIndex]
        //     })
        // );

        _onEventImport(
            EVMEventInfo({
                blockchainID: bytes32(0),
                blockNumber: 0,
                txIndex: 0,
                logIndex: 0,
                log: EVMLog({loggerAddress: address(0), topics: new bytes32[](0), data: bytes("")})
            })
        );
    }

    function _onEventImport(EVMEventInfo memory) internal {
        emit WeGotIt();
    }
}
