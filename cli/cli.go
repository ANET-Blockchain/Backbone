package cli

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct{}

func (cli *CLI) PrintUsage() {
	fmt.Println("Usage:")
	fmt.Println("  CreateBlockchain -address [ADDRESS] : Create a blockchain and send genesisblock reward to ADDRESS")
	fmt.Println("  CreateWallet : Generates a new key-pair and saves it into the wallet file")
	fmt.Println("  GetBalance -address [ADDRESS] : GET balance of ADDRESS")
	fmt.Println("  GetAllBalances : print the balances of all addresses")
	fmt.Println("  ListAddresses : Lists all addresses from the wallet file")
	fmt.Println("  PrintChain : print all the blocks of the blockchain")
	fmt.Println("  Send -from [FROM ADDRESS] -to [TO ADDRESS] -amount [AMOUNT] : Send AMOUNT of coins from FROM to TO. Mine on the same node, when -mine is set.")
	fmt.Println("  StartNode -miner [ADDRESS] : Start a node with ID specified in NODE_ID env. var. -miner enables mining")
	fmt.Println("  ReindexUTXO : Rebuilds the UTXO set")
}

func (cli *CLI) ValidateArgs() {
	if len(os.Args) < 2 {
		cli.PrintUsage()
		os.Exit(1)
	}
}

func (cli *CLI) Run() {
	cli.ValidateArgs()

	nodeID := os.Getenv("NODE_ID")
	if nodeID == "" {
		fmt.Println("NODE_ID env. var is not set!")
		os.Exit(1)
	}

	getBalanceCmd := flag.NewFlagSet("GetBalance", flag.ExitOnError)
	getAllBalancesCmd := flag.NewFlagSet("GetAllBalances", flag.ExitOnError)
	createBlockchainCmd := flag.NewFlagSet("CreateBlockchain", flag.ExitOnError)
	createWalletCmd := flag.NewFlagSet("CreateWallet", flag.ExitOnError)
	listAddressesCmd := flag.NewFlagSet("ListAddresses", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("Send", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("PrintChain", flag.ExitOnError)
	startNodeCmd := flag.NewFlagSet("StartNode", flag.ExitOnError)
	reindexUTXOcmd := flag.NewFlagSet("ReindexUTXO", flag.ExitOnError)

	getBalanceAddress := getBalanceCmd.String("address", "", "The address to get balance for")
	createBlockchainAddress := createBlockchainCmd.String("address", "", "The address to send genesis block reward to")
	sendFrom := sendCmd.String("from", "", "Source wallet address")
	sendTo := sendCmd.String("to", "", "Destination wallet address")
	sendAmount := sendCmd.Int("amount", 0, "Amount to send")
	sendMine := sendCmd.Bool("mine", false, "Mine immediately on the same node")
	startNodeMiner := startNodeCmd.String("miner", "", "Enable mining mode and send reward to ADDRESS")

	switch os.Args[1] {
	case "GetBalance":
		err := getBalanceCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "GetAllBalances":
		err := getAllBalancesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "CreateBlockchain":
		err := createBlockchainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "CreateWallet":
		err := createWalletCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "ListAddresses":
		err := listAddressesCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "PrintChain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "Send":
		err := sendCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "ReindexUTXO":
		err := reindexUTXOcmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "StartNode":
		err := startNodeCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		cli.PrintUsage()
		os.Exit(1)
	}

	if getBalanceCmd.Parsed() {
		if *getBalanceAddress == "" {
			getBalanceCmd.Usage()
			os.Exit(1)
		}
		cli.GetBalance(*getBalanceAddress, nodeID)
	}

	if getAllBalancesCmd.Parsed() {
		cli.GetAllBalances(nodeID)
	}

	if createBlockchainCmd.Parsed() {
		if *createBlockchainAddress == "" {
			createBlockchainCmd.Usage()
			os.Exit(1)
		}
		cli.CreateBlockchain(*createBlockchainAddress, nodeID)
	}

	if createWalletCmd.Parsed() {
		cli.CreateWallet(nodeID)
	}

	if listAddressesCmd.Parsed() {
		cli.ListAddresses(nodeID)
	}

	if printChainCmd.Parsed() {
		cli.PrintChain(nodeID)
	}

	if sendCmd.Parsed() {
		if *sendFrom == "" || *sendTo == "" || *sendAmount <= 0 {
			sendCmd.Usage()
			os.Exit(1)
		}

		cli.Send(*sendFrom, *sendTo, *sendAmount, nodeID, *sendMine)
	}

	if startNodeCmd.Parsed() {
		nodeID := os.Getenv("NODE_ID")
		if nodeID == "" {
			startNodeCmd.Usage()
			os.Exit(1)
		}
		cli.startNode(nodeID, *startNodeMiner)
	}

	if reindexUTXOcmd.Parsed() {
		cli.ReindexUTXO(nodeID)
	}
}
