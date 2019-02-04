package src

import (
	"errors"
)

import "github.com/boltdb/bolt"

type BlockchainIterator struct {
	currentHash	[]byte
	db		*bolt.DB
}

/* Get the next block in the chain using iterator's currentHash */
func (itr *BlockchainIterator) Next() *Block {
	var block *Block

	err := itr.Db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blocksBucket))
		encodedBlock := bkt.Get(itr.currentHash)
		block = Deserialize(encodedBlock)
		/* Handle the End of File exception */
		if block == nil {
			err := errors.New(EOFerr)
			return err
		}
		return nil
	})
	/* This should happen only when the blockchain is finished */
	if err != nil {
		return nil
		//log.Fatal("Error in viewing the DB file")
	}

	/* Set the pointer for iteration of previous block */
	itr.currentHash = block.PrevBlockHash
	return block
}
