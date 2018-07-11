package main

import "github.com/boltdb/bolt"

type Blockchain struct {
	tip		[]byte
	db		*bolt.DB
}

/* Add a new block to existing blockchain */
func (bc *Blockchain) Addblock(data string) {
	// prevBlock := bc.blocks[len(bc.blocks)-1]
	// newBlock := NewBlock(data, prevBlock.Hash)
	// bc.blocks = append(bc.blocks, newBlock)

	var lastHash []byte
	/* Retrieve ? hash */
	err := bc.db.View(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blocksBucket))
		lastHash = bkt.Get([]byte("1"))
		return nil
	})

	newBlock := NewBlock(data, lastHash)
	err := bc.db.Update(func(tx *bolt.Tx) error {
		bkt := tx.Bucket([]byte(blocksBucket))
		err := bkt.Put(newBlock.Hash, newBlock.Serialize())
		// if err != nil {
		// 	log.Fatal("Couldn't store the genesis block")
		// }
		err := bkt.Put([]byte("1"), newBlock.Hash)
		// if err != nil {
		// 	log.Fatal("Couldn't store the genesis hash")
		// }
		bc.tip = newBlock.Hash
	})
	return nil
}

/* Create and return an iterator for the blockchain */
func (bc *Blockchain) Iterator() *BlockchainIterator {
	bcitr := &BlockchainIterator{bc.tip, bc.db}
	return bcitr
}
