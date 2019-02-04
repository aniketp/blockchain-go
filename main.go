package main

import (
	"github.com/aniketp/blockchain-go/src"
)

func main() {
	bc := src.NewBlockChain()

	// bc.Addblock("Send 1 BTC to Lavannya")
	// bc.Addblock("Send 2 more BTC to Lavannya")
	// bc.Addblock("Send 1 Heart to Apoorva")
	// bc.Addblock("Send 2 more Hearts to Apoorva")
	// bc.Addblock("Send 1 more BTC to Lavannya")
	//
	// for _, block := range bc.blocks {
	// 	pow := NewProofOfWork(block)
	// 	fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
	// 	fmt.Printf("Data: %s\n", block.Data)
	// 	fmt.Printf("Hash: %x\n", block.Hash)
	// 	fmt.Printf("Valid PoW: %s\n",
	// 		strconv.FormatBool(pow.Validate()))
	// 	fmt.Println()
	// }

	defer bc.db.Close()
	cli := src.CLI{bc}
	cli.Run()
}
