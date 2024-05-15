// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {IEventImporter} from "./IEventImporter.sol";

abstract contract EventImporter is IEventImporter {
    /*
     * @notice 
     * 1. Imports a block hash from another blockchain via Warp.
     * 2. Verifies that the provided blockHeader matches the authenticated block hash.
     * 3. Gets the receipt at the given transaction index by verifying the merkle proof against the block header's receipt root.
     * 4. Decodes and returns the log at the given log index from the receipt.
     */
    function importEvent(bytes calldata blockHeader, uint256 txIndex, bytes calldata receiptProof, uint256 logIndex)
        external
        returns (EVMEventInfo memory);
}
