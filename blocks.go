package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
)

type Block struct {
	Timestamp  	int64		/* Time of block creation */
	Transactions	[]*Transaction	/* Valuable info in the block */
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

/* Serialize the blocks before storing as key value pairs */
func (b *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)

	err := encoder.Encode(b)
	if err != nil {
		panic("Encoding failed")
	}
	return result.Bytes()
}

/* Hash the transactions to be used while preparing hash of block */
func (b *Block) HashTransactions() []byte {
	var txHashes [][]byte
	var txHash [32]byte

	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.ID)
	}
	txHash = sha256.Sum256(bytes.Join(txHashes, []byte{}))
	return txHash[:]
}

/* A common interface for deserializing any block value */
func Deserialize (d []byte) * Block {
	var block Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil
	}
	return &block
}
