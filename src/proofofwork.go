package main

import (
	"bytes"
	"crypto/sha256"
	"log"
	"fmt"
	"math"
	"math/big"
)

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

/* Validate that the computed SHA satisfies the criteria */
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data := pow.prepareData(pow.block.Nonce)
	hash := sha256.Sum256(data)
	hashInt.SetBytes(hash[:])

	return (hashInt.Cmp(pow.target) < 0)
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

		/* If the calculated hash was less than target then we exit */
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
