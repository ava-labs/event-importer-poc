# Cross-Chain EVM Event Importing

> **TL;DR: This repo is a proof-of-concept for how to import arbitrary EVM event logs emitted by chains on other Subnets. It uses Avalanche Warp Messaging to authenticate a block hash from a chain in another Subnet, and Merkle proofs to validate that the event to be imported was emitted in that block.**

[Avalanche Warp Messaging](https://docs.avax.network/build/cross-chain/awm/overview) (AWM) allows for Subnets to validate that arbitrary messages have been signed by a configurable threshold of current stake weight of other Subnets. Currently, the most common type of cross-chain interaction involves a source chain transaction which sends a message, and a destination chain transaction which delivers the message. In the case of AWM, the source chain transaction would call the Warp precompile `sendWarpMessage` interface function, and the destination chain transaction would call the Warp precompile `getVerifiedWarpMessage` interface function.

The Warp precompile also offers a `getVerifiedWarpBlockHash` interface function, which surfaces 32-byte hashes that were signed by a threshold of stake weight of a given Subnet. Validators of EVM chains that support Warp should be willing to sign the hash of any accepted block on that chain. **This enables importing information about that chain into other Subnets without needing to interact with the source chain at all.**

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
- Can/Should block hash authenticated via Warp messages be stored such that they can be used to prove arbitrarily many events going forward without having to re-verify a Warp aggregate signature?
- Is the Merkle proof verification or RLP decoding prohibitively expensive gas-wise? Can they be optimized if so? How does gas usage scale with the number of receipts in the block including the event to be imported?
- Is the delay from the time an event is emitted on a source chain to when it is imported on another chain a non-starter for certain applications? 
    - For instance, when importing a price feed stream, it is known what the next value to be imported will be before it is actually imported.
    - If yes, what is the acceptable delay, or how can the risk be mitigated? Note that some delay already exists in the single chain case from when the transaction hits the mempool of the source chain to when it gets included in a block.
- How can relayers be incentivized to deliver new events that meet certain criteria?
- What is the preferred mechanism for delivering events to be imported on other chains?
    - Could take the form of a relayer application that listens for certain events and sends them along with the required proof to be imported to pre-configured chains and contracts.
    - Could take the form of a public "get Warp block signature" API, that allows UIs to construct transactions importing events from other chains. In this model, users would import their own events from their wallet. The API service could potentially pre-emptively construct the aggregate signature for each block on the set of supported chains such that it can serve them on request without needing to query validators for their individual BLS signatures.
