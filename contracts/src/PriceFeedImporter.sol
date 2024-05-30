// SPDX-License-Identifier: Ecosystem

pragma solidity ^0.8.17;

import {EVMEventInfo, EventImporter} from "./EventImporter.sol";

/**
 * @notice An example EventImporter implementation that imports the latest price feed data from another blockchain.
 */
contract PriceFeedImporter is EventImporter {
    bytes32 public constant ANSWER_UPDATED_EVENT_SIGNATURE = keccak256("AnswerUpdated(int256,uint256,uint256)");

    // Blockchain ID of the oracle chain.
    bytes32 public immutable sourceBlockchainID;

    // Address of the Aggregator contract on the source blockchain.
    address public immutable sourceOracleAggregator;

    // Latest answer information.
    int256 public currentAnswer;
    uint80 public roundID;
    uint256 public updatedAt;

    // The block and transaction on the source blockchain where the latest answer was updated.
    uint256 public latestSourceBlockNumber;
    uint256 public latestSourceTxIndex;
    uint256 public latestSourceLogIndex;

    // Event emitted when the answer is updated.
    event AnswerUpdated(int256 currentAnswer, uint80 roundID, uint256 updatedAt);

    constructor(bytes32 sourceBlockchainID_, address sourceOracleAggregator_) {
        sourceBlockchainID = sourceBlockchainID_;
        sourceOracleAggregator = sourceOracleAggregator_;
    }

    /**
     * @notice Returns the latest round data if available.
     */
    function latestRoundData() external view returns (uint80, int256, uint256, uint256, uint80) {
        require(updatedAt != 0, "No data");
        return (roundID, currentAnswer, updatedAt, updatedAt, roundID);
    }

    function _onEventImport(EVMEventInfo memory eventInfo)
        internal
        override
        _onlySourceOracleEvents(eventInfo)
        _onlyValidAnswerUpdatedEvents(eventInfo)
        _onlyMoreRecentEvents(eventInfo)
    {
        // Update the latest answer.
        currentAnswer = int256(uint256(eventInfo.log.topics[1]));
        roundID = uint80(uint256(eventInfo.log.topics[2]));
        updatedAt = uint256(bytes32(eventInfo.log.data));

        // Update the latest source block information.
        latestSourceBlockNumber = eventInfo.blockNumber;
        latestSourceTxIndex = eventInfo.txIndex;
        latestSourceLogIndex = eventInfo.logIndex;

        emit AnswerUpdated(currentAnswer, roundID, updatedAt);
    }

    modifier _onlySourceOracleEvents(EVMEventInfo memory eventInfo) {
        require(eventInfo.blockchainID == sourceBlockchainID, "Invalid blockchain ID");
        require(eventInfo.log.loggerAddress == sourceOracleAggregator, "Invalid logger address");
        _;
    }

    modifier _onlyValidAnswerUpdatedEvents(EVMEventInfo memory eventInfo) {
        require(eventInfo.log.topics.length == 3, "Invalid log topics");
        require(eventInfo.log.data.length == 32, "Invalid log data");
        require(eventInfo.log.topics[0] == ANSWER_UPDATED_EVENT_SIGNATURE, "Invalid event signature");
        _;
    }

    modifier _onlyMoreRecentEvents(EVMEventInfo memory eventInfo) {
        require(
            eventInfo.blockNumber >= latestSourceBlockNumber
                && (eventInfo.blockNumber > latestSourceBlockNumber || eventInfo.txIndex > latestSourceTxIndex)
                && (
                    eventInfo.blockNumber > latestSourceBlockNumber || eventInfo.txIndex > latestSourceTxIndex
                        || eventInfo.logIndex > latestSourceLogIndex
                ),
            "Stale event"
        );
        _;
    }
}
