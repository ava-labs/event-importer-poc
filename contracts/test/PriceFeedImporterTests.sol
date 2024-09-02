// (c) 2024, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

// SPDX-License-Identifier: Ecosystem

pragma solidity 0.8.18;

import {Test} from "forge-std/Test.sol";
import {WarpBlockHash, IWarpMessenger} from "@subnet-evm/contracts/interfaces/IWarpMessenger.sol";
import {PriceFeedImporter} from "src/PriceFeedImporter.sol";
import {EVMEventInfo, EVMLog} from "src/IEventImporter.sol";

contract PriceFeedImporterMock is PriceFeedImporter(bytes32(0), address(0), 0, "", 0) {
    function importEvent(EVMEventInfo memory eventInfo) external {
        _onEventImport(eventInfo);
    }
}

contract PriceFeedImporterTest is Test {
    address public constant WARP_PRECOMPILE_ADDRESS = 0x0200000000000000000000000000000000000005;
    bytes32 public constant DEFAULT_BLOCKCHAIN_ID =
        bytes32(hex"55e1fcfdde01f9f6d4c16fa2ed89ce65a8669120a86f321eef121891cab61241");
    address public constant SOURCE_ORACLE_AGGREGATOR = 0x154baB1FC1D87fF641EeD0E9Bc0f8a50D880D2B6;

    PriceFeedImporterMock public priceFeedImporter;

    event AnswerUpdated(int256 currentAnswer, uint80 roundID, uint256 updatedAt);
    event EventImported(
        bytes32 indexed sourceBlockchainID,
        bytes32 indexed sourceBlockHash,
        address indexed loggerAddress,
        uint256 txIndex,
        uint256 logIndex
    );

    function setUp() public virtual {
        priceFeedImporter = new PriceFeedImporterMock();
    }

    function testSetPriceFeed() public {
        int256 answer = 123456;
        uint80 logRoundID = 0;
        uint256 logUpdatedAt = 1725274805;

        EVMEventInfo memory eventInfo = buildEVMEvent(answer, logRoundID, logUpdatedAt);
        priceFeedImporter.importEvent(eventInfo);

        (uint80 roundID, int256 currentAnswer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound) =
            priceFeedImporter.latestRoundData();
        assertEq(roundID, logRoundID);
        assertEq(currentAnswer, answer);
        assertEq(startedAt, logUpdatedAt);
        assertEq(updatedAt, logUpdatedAt);
        assertEq(answeredInRound, logRoundID);
    }

    function testIncreasingRoundID() public {
        priceFeedImporter.importEvent(buildEVMEvent(0, 0, 1));
        priceFeedImporter.importEvent(buildEVMEvent(0, 1, 1));
        vm.revertWith("roundID should be monotonically increasing");
        priceFeedImporter.importEvent(buildEVMEvent(0, 0, 1));
    }

    function testHistoricalRoundData() public {
        priceFeedImporter.importEvent(buildEVMEvent(0, 0, 1));
        priceFeedImporter.importEvent(buildEVMEvent(0, 1, 1));
        vm.revertWith("roundID should be monotonically increasing");
        priceFeedImporter.importEvent(buildEVMEvent(0, 0, 1));
    }

    function buildEVMEvent(int256 answer, uint80 roundID, uint256 updatedAt) private view returns (EVMEventInfo memory) {
        bytes32[] memory topics = new bytes32[](3);
        topics[0] = priceFeedImporter.ANSWER_UPDATED_EVENT_SIGNATURE();
        topics[1] = bytes32(uint256(answer));
        topics[2] = bytes32(uint256(uint80(roundID)));

        EVMLog memory log = EVMLog({
            loggerAddress: address(0),
            topics: topics,
            data: abi.encode(bytes32(updatedAt))
        });

        return EVMEventInfo({
            blockchainID: bytes32(0),
            blockNumber: 1,
            txIndex: 0,
            logIndex: 0,
            log: log
        });
    }
}
