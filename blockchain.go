package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp  	int64		/* Time of block creation */
	Data		[]byte		/* Valuable info in the block */
	PrevBlockHash	[]byte		/* Hash of the previous block */
	Hash 		[]byte		/* Hash of the block */
}

type Blockchain struct {
	blocks []*Block
}

/* Calculate the Hash value of current block using it's metadata */
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	/* Concatenate Data, PrevBlockHash & Hash together */
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data,
		timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	/* Set the computed hash as the block's hash value */
	b.Hash = hash[:]
}

/* Create and return a new block */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data),
		prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

/* Add a new block to existing blockchain */
func (bc *Blockchain) Addblock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	bc := NewBlockChain()

	bc.Addblock("Send 1 BTC to Lavannya")
	bc.Addblock("Send 2 more BTC to Lavannya")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
