# Cross-Chain EVM Event Importing

> **TL;DR: This repo is a proof-of-concept for how to import arbitrary EVM event logs emitted by chains on other Subnets. It uses Avalanche Warp Messaging to authenticate a block hash from a chain in another Subnet, and Merkle proofs to validate that the event to be imported was emitted in that block.**

[Avalanche Warp Messaging](https://docs.avax.network/build/cross-chain/awm/overview) (AWM) allows for Subnets to validate that arbitrary messages have been signed by a configurable threshold of current stake weight of other Subnets. Currently, the most common type of cross-chain interaction involves a source chain transaction which sends a message, and a destination chain transaction which delivers the message. In the case of AWM, the source chain transaction would call the Warp precompile `sendWarpMessage` interface function, and the destination chain transaction would call the Warp precompile `getVerifiedWarpMessage` interface function.

The Warp precompile also offers a `getVerifiedWarpBlockHash` interface function, which surfaces 32-byte hashes that were signed by a threshold of stake weight of a given Subnet. Validators of EVM chains that support Warp should be willing to sign the hash of any accepted block on that chain. **This enables importing information about that chain into other Subnets without needing to interact with the source chain at all.**

In order to make use of an authenticated block hash to import event logs, the block header must be provided that corresponds to the given block hash. The block header includes a `receiptsRoot`, which is the root hash of a Merkle tree constructed with the receipt for every transaction in the block in order. A Merkle proof can be provided against this root to proof the inclusion of a specific receipt in the block. Transaction receipts include each of the events emitted as a part of that transaction.

## Example Implementations
Contracts looking to import events from other chains should inherit from the abstract `EventImporter` contract in this repo. The `EventImporter` contract handles the authentication of block headers, verification of Merkle proofs, and parsing of event logs. The child contracts import events should provide their own `_onEventImport` implementation to define the logic that should be executed when a new event is received. For example:
```solidity
pragma solidity 0.8.18;

import {EVMEventInfo, EventImporter} from "./EventImporter.sol";

contract MyEventImporter is EventImporter {
    // Blockchain ID of the source chain.
    bytes32 public immutable sourceBlockchainID;

    // Address of contract to import events from on the source chain.
    address public immutable emittingContract;

    constructor(bytes32 sourceBlockchainID_, address emittingContract_) {
        sourceBlockchainID = sourceBlockchainID_;
        emittingContract = emittingContract_;
    }

    function _onEventImport(EVMEventInfo memory eventInfo)
        internal
        override
    {
        require(eventInfo.blockchainID == sourceBlockchainID);
        require(eventInfo.log.loggerAddress == emitterContract)

        // Custom logic here
        // ....
    }
}

```

## Structure
- `contracts/` is a Foundry project that implments the necessary Merkle proof verification and RLP decoding to be able to import event logs from other Subnets.
    - `EventImporter.sol` is an abstract contract that handles the Warp verification, Merkle proof verification, and RLP decoding for a given event against a provided block hash and header. Child contracts looking to import events only have to implement the `_onEventImport` function to define how they would like to handle specific event imports.
    - `PriceFeedImporter.sol` is an example event importer contract that demonstrates how one might import updates to a price feed from another Subnet.
- `proofs/` contains Golang utilies for constructing Merkle proofs of inclusion of given transaction receipts from blocks.
    - `proof_utils.go` contains helper function for constructing Merkle proof to be used by other Golang applications (including E2E tests).
    - `tool/proof_build.go` contains a program to fetch a block header and construct a Merkle proof for a specific transaction within that block. The proof is printed in the logs and can be used in unit tests or for debugging.
        - Exmaple usage: `go run proofs/tools/proof_builder.go --rpc-url https://api.avax.network/ext/bc/C/rpc --block-hash 0x27bfa26de1022c4d9cbbc67d6819037ce8a4de99832f00105d7be9d2b4d25369 --tx-index 1`
- `tests/` includes Ginkgo end-to-end test of the Solidity contracts, serving as reference implementation for constructing transactions that successfully import event logs from other Subnets.
- `abi-bindings/` includes Go ABI bindings for the contracts in contracts/.

## Test it out
To run the Solidity unit tests:
```
cd contracts
forge test -vvv
```
To run the E2E tests:
```
./scripts/e2e_test.sh
```
The E2E test flow:
1. Starts a local Avalanche network with a Subnet.
2. Deploys a "mock price feed" contract to the C-Chain
3. Deploys a `PriceFeedImporter` contract on the Subnet. 
4. Sends a transaction to update the price feed value on the C-Chain, as would be done by an oracle provider. 
5. Imports the price feed update event into the Subnet by constructing a Warp signature of the block hash that it occured in on the C-Chain, constructing a Merkle proof for the transaction receipt the event was contained in, and broadcasting this information to the Subnet by calling the `importEvent` interface function of the `PriceFeedImporter` contract.

## Open Questions and Considerations
### Should block hashes authenticated via Warp messages be stored such that they can be used to prove arbitrarily many events going forward without having to re-verify a Warp aggregate signature?

The `IEventImporter` interface definition intentionally includes an explicit `sourceBlockchain` parameter such that implementations do not necessarily need to use the Warp precompile to directly authentication block hashes. Instead, previously authenticated block hashes can be checked against an external "block hash registry"-like contract. This would enable multiple events to be imported from the same block with only requiring the block hash be authenticated once.

### Is the Merkle proof verification or RLP decoding prohibitively expensive gas-wise? Can they be optimized if so? How does gas usage scale with the number of receipts in the block including the event to be imported?

Currently, roughly ~450,000 gas is used in transactions in the E2E. These transaction include:
- Verifying a Warp signature with 4 signers
- Verifying a Merkle proof of inclusion of a receipt in a block that contains a single transaction

Additional signers cost [500 gas each](https://github.com/ava-labs/subnet-evm/blob/master/precompile/contracts/warp/contract.go#L34). Further analysis is needed to determine how gas usage increases as the number of transactions in blocks and number of logs in the relevant transaction grows.

### Is the delay from the time an event is emitted on a source chain to when it is imported on another chain a non-starter for certain applications? 

For potential applications such as importing data feeds from other chains, it is known what the next value to be imported should be before the import actually occurs. This could present an MEV-like opportunity for applications that depend on the data feed values. Note that some level of delay already exists in the single chain case from when a transaction hits the mempool of the source chain to when it gets included in a block on that chain. If this increased delay is a concern, what amount of delay is acceptable, and how can that risk be mitigated?

### How can relayers be incentivized to deliver new events that meet certain criteria?

Contracts importing events could theoretically define rewards able to be claimed by anyone that delivers events that meet specific criteria. 

### What are the preferred mechanisms for delivering events to be imported on other chains?

Possible options include:
- A relayer application that listens for specific events and sends them along with the required proof to be imported to pre-configured chains and contracts.
- A public "get Warp block signature" API that allows UIs to construct transactions that import events. In this model, users would import their own events from their wallet. One implementation of this [API already exists with subnet-evm](https://github.com/ava-labs/subnet-evm/blob/master/warp/service.go#L80), but is not publicly available since it is a potential DOS vector for nodes that have it enabled.
