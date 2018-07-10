package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"time"
)

var targetBits int64 = 18		/* Increasing this would frustate CPU */

type Block struct {
	Timestamp  	int64		/* Time of block creation */
	Data		[]byte		/* Valuable info in the block */
	PrevBlockHash	[]byte		/* Hash of the previous block */
	Hash 		[]byte		/* Hash of the block */
	Nonce		int
}

type Blockchain struct {
	blocks		[]*Block
}

type ProofOfWork struct {
	block 		*Block
	target		*big.Int
}

/* Prepare the data to be hashed */
func (pow *ProofOfWork) prepareData(nonce int) []byte {
	data := bytes.Join([][]byte{
		pow.block.PrevBlockHash,
		pow.block.Data,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	}, []byte{})
	return data
}

/* Implementation of Core Proof of Work algorithm */
func (pow *ProofOfWork) Run() (int , []byte) {
	var hashInt big.Int
	var hash [32]byte
	maxNonce := math.MaxInt64
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < maxNonce {
		/* Retrieve the data for given PoW */
		data := pow.prepareData(nonce)
		/* Calculate it's SHA hash */
		hash := sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		/* If the calculated hash was less than target then we  */
		if hashInt.Cmp(pow.target) == -1 {
			fmt.Printf("\n\n")
			return nonce, hash[:]
		} else {
			nonce++
		}
	}
	/* Should not have reached here: nonce > maxNonce */
	log.Fatal("Couldn't compute the allowed hash in time-limit")
	return nonce, hash[:]	/* Placeholder return value */
}

/* Calculate New proof of work for a corresponding block */
func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	pow := &ProofOfWork{b, target}
	return pow
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
		prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	// block.SetHash()
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
	bc.Addblock("Send 1 Heart to Apoorva")
	bc.Addblock("Send 2 more Hearts to Apoorva")
	bc.Addblock("Send 1 more BTC to Lavannya")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}

func IntToHex(n int64) []byte {
    return []byte(strconv.FormatInt(n, 16))
}
