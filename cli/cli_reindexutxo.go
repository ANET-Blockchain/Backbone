package cli

import (
	"Vanilla/core"
	"fmt"
)

func (cli *CLI) ReindexUTXO(nodeID string) {
	bc := core.NewBlockchain(nodeID)
	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	count := UTXOSet.CountTransactions()
	fmt.Printf("Done! There are %d transactions in tohe UTXO set.\n", count)
}
