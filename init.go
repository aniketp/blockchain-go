package main

import (
	"math/big"
	"time"
)

/* Create and return a new block */
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data),
		prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	// block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBlockChain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

/* Calculate New proof of work for a corresponding block */
func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}
