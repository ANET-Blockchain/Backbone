package cli

import (
	"Vanilla/core"
	"fmt"
	"log"
)

func (cli *CLI) ListAddresses(nodeID string) {
	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}

}

func (cli *CLI) GetAllBalances(nodeID string) {
	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		cli.GetBalance(address, nodeID)
	}

}
