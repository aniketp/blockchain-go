package main

import "github.com/boltdb/bolt"

type BlockchainIterator struct {
	currentHash	[]byte
	db		*bolt.DB
}

/* Get the next block in the chain using iterator's currentHash */
func (itr *BlockchainIterator) Next() *Block {
	var block *Block

	err := itr.db.View(func{tx *bolt.Tx} error {
		bkt := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bkt.Get(itr.currentHash)
		block = Deserialize(encodedBlock)
		return nil
	})

	/* Set the pointer for iteration of previous block */
	itr.currentHash = block.PrevBlockHash
	return block
}
