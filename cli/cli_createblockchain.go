package cli

import (
	"Backbone/core"
	"fmt"
	"log"
)

func (cli *CLI) CreateBlockchain(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR : Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.Db.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
