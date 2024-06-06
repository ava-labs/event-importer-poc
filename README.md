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

# High Level Project Plan
This repo is a proof-of-concept of the idea described above. In order to turn it into a production ready framework it needs to be expanded and made more robust.

## Objective and Use Cases
It should be incredibly easy for a new Subnet to securely receive event emissions and state changes from contracts on other chains. Unlike the `TeleporterMessenger` contracts which provide point-to-point messaging from one chain to another to be used in building cross-chain applications, this primitive allows for arbitrary data to be imported into a Subnet from other chains without any interaction required on the other chains.

This could be used to:
- Prove an address holds an NFT/NTT issued on other chains
    - Could be used to give special access within games to holders of rare badges/coupons on other chains
    - Could be used to build Avalanche-based ID/credentials
    - etc
- Import event feeds from other chains such as price feeds.

## Open Questions by Component
### Solidity Contract Utilities

#### Requirements 
The Solidity utility contract must support:
- Receiving Warp block hash messages.
- Taking an RLP encoded block header and verifying that it matches a block hash authenticated using Warp.
- Verifying a Merkle proof of an RLP encoded receipt against the receipt root of a block header.
- Verifying a Merkle proof of a contracts state against the state root of a block header.
- RLP decoding of block headers and receipts.

#### Open Questions and Considerations
- Can/Should block hashes authenticated via Warp messaged be stored such that they can be used to prove arbitrarily many events going forward without having to re-verify a Warp aggregate signature?
- Is the Merkle proof verification or RLP decoding too expensive gas-wise, and can they be optimized?
- What frameworks should we have to decide which events are valid to be imported by an application? (i.e. strict ordering, most recent only, archival, etc)
- Should we add a mechanism to define incentives for deliverers of events that meet certain criteria? Could we safely use the native minter precompile to guarantee profitability of deliverers (i.e. refund the transaction fee + some additional amount)?

### Publisher Service and/or AWM Relayer Extension
#### Requirements
There should be an application that can be configured and run out-of-the-box to publish specified events to pre-configured event receiver contracts. There should be an API made available for UIs to fetch an aggregate signature of a block hash itself, such that users can easily send their own transactions to import events.

#### Open Questions and Considerations
- How can we ensure that a "get Warp signature" public API is not a DOS vector? Pre-emptively generate and cache signatures for blocks to serve on request?
- Should a service to publish pre-configured events be embedded within the AWM relayer itself? Or should it be its own service/sidecar that is pointed to an API to be able to fetch Warp signatures when needed?
- What are the necessary configuration settings?
    - How to know which events to publish?
    - How to know where to publish them to?
    - How to check for profitability of publishing an event?




