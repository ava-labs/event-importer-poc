package main

import (
	"encoding/json"
	"os"

	"github.com/ava-labs/subnet-evm/core/types"
	"github.com/ava-labs/subnet-evm/trie"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

func main() {
	log.Root().SetHandler(log.StreamHandler(os.Stdout, log.LogfmtFormat()))
	receipt1JSON := `{
        "blockHash": "0x3c490d675f4c0fb921a183cb093748f8a70fbca703695d9cec0da15809020169",
        "blockNumber": "0x1eea71b",
        "contractAddress": null,
        "cumulativeGasUsed": "0x1ac55",
        "effectiveGasPrice": "0x6f332da80",
        "from": "0x5106af71713d3be15415107f61f6bc195ccabcce",
        "gasUsed": "0x1ac55",
        "logs": [
            {
                "address": "0xa9d587a00a31a52ed70d6026794a8fc5e2f5dcb0",
                "topics": [
                    "0x43dc749a04ac8fb825cbd514f7c0e13f13bc6f2ee66043b76629d51776cff8e0",
                    "0x0000000000000000000000000000000000000000000000000000000000001bda"
                ],
                "data": "0x000000000000000000000000080c0a5c7369739e298a8709eda8924941e0f77d",
                "blockNumber": "0x1eea71b",
                "transactionHash": "0xba9622e1c05563cdb356ebebb0877a3ab0efe57ffade32bb1e92841e66a4ab95",
                "transactionIndex": "0x0",
                "blockHash": "0x3c490d675f4c0fb921a183cb093748f8a70fbca703695d9cec0da15809020169",
                "logIndex": "0x0",
                "removed": false
            }
        ],
        "logsBloom": "0x00000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080000000000000000100000000000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000080000000000000000000000000000000000000000000000000000000000000000000000000000000000010000000021000000000000000000000000000000000000000000000000000000000000000000000000000",
        "status": "0x1",
        "to": "0xa9d587a00a31a52ed70d6026794a8fc5e2f5dcb0",
        "transactionHash": "0xba9622e1c05563cdb356ebebb0877a3ab0efe57ffade32bb1e92841e66a4ab95",
        "transactionIndex": "0x0",
        "type": "0x2"
    }`
	var receipt1 types.Receipt
	err := json.Unmarshal([]byte(receipt1JSON), &receipt1)
	if err != nil {
		panic(err)
	}
	log.Info("Unmarshalled receipt", "receipt", receipt1)

	t, err := trie.New(trie.StateTrieID(common.Hash{}), trie.NewDatabase(nil))
	if err != nil {
		panic(err)
	}
	receipts := []*types.Receipt{&receipt1}
	root := types.DeriveSha(types.Receipts(receipts), t)
	log.Info("Computed trie root", "root", root.String())

}
