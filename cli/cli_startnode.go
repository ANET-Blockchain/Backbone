package cli

import (
	"Backbone/core"
	"fmt"
	"log"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if core.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards : ", minerAddress)
		} else {
			log.Panic("Wrong miner Address!")
		}
	}

	// start http server simultaneously
	go core.HttpServer()

	// blockchain node start
	core.StartServer(nodeID, minerAddress)
}
