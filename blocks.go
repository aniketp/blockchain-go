package main

import (
	"bytes"
	"crypto/sha256"
	"strconv"
)

type Block struct {
	Timestamp  	int64		/* Time of block creation */
	Data		[]byte		/* Valuable info in the block */
	PrevBlockHash	[]byte		/* Hash of the previous block */
	Hash 		[]byte		/* Hash of the block */
	Nonce		int		/* Random nonce for PoW */
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
