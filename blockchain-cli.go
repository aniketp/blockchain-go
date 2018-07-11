package main

import (
	"log"
	"flag"
	"fmt"
	"strconv"
	"os"
)

type CLI struct {
	bc *Blockchain
}

/* Driver function for the blockchain's CLI */
func (cli *CLI) Run() {
	cli.validateArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	addBlockData := addBlockCmd.String("data", "", "Block data")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("Error in parsing arguments")
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("Error in parsing arguments")
		}
	default:
		cli.printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *addBlockData == "" {
			/* No additional arguments were supplied */
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}
}

/* Following are the utility functions to do what the user desired */
func (cli *CLI) addBlock(data string) {
	cli.bc.Addblock(data)
	fmt.Println("Success!")
}

func (cli *CLI) printChain() {
	bci := cli.bc.Iterator()

	for {
		/* Iterate over blocks using the custom built iterator */
		block := bci.Next()
		if block == nil {
			break
		}
		pow := NewProofOfWork(block)

		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
	}
}

func (cli *CLI) validateArgs() {
	if os.Args[0] != "blockchain" {
		cli.printUsage()
		os.Exit(1)
	}
	if !(os.Args[1] == "addblock" || os.Args[1] == "printchain") {
		cli.printUsage()
		os.Exit(1)
	}
}

func (cli *CLI) printUsage() {
	fmt.Printf("Usage: blockchain [addblock] [printchain ...]")
}
