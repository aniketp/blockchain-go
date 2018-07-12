package main

import (
	"log"
	"math/big"
	"time"

	"github.com/boltdb/bolt"
)

/* Create and return a new block */
func NewBlock(transactions []*Transaction, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), transactions,
		prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash
	block.Nonce = nonce
	// block.SetHash()
	return block
}

/* Return the new (first) genesis block */
func NewGenesisBlock(coinbase *Transaction) *Block {
	return NewBlock([]*Transaction{coinbase}, []byte{})
}

/* This function checks for the existence of a bucket in the database file, */
/* if it exists, retrieves the value else creates one with the genesis hash */
func NewBlockChain(address string) *Blockchain {
	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	if err != nil {
		log.Fatal("Couldn't open the database")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		cbtx := NewCoinbaseTX(address, genesisCoinbaseData)
		b := tx.Bucket([]byte(blocksBucket))

		/* The Bucket does not exist, hence create one */
		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(blocksBucket))
			if err != nil {
				log.Fatal("Couldn't create a new bucket")
			}
			err = b.Put(genesis.Hash, genesis.Serialize())
			// if err != nil {
			// 	log.Fatal("Couldn't store the genesis block")
			// }
			err = b.Put([]byte("1"), genesis.Hash)
			// if err != nil {
			// 	log.Fatal("Couldn't write to the genesis block")
			// }
			tip = genesis.Hash
		} else {
			/* Since the bucket exists, retrieve first value */
			tip = b.Get([]byte("1"))
		}
		return err
	})

	bc := Blockchain{tip, db}
	return &bc
	//return &Blockchain{[]*Block{NewGenesisBlock()}}
}

/* Calculate New proof of work for a corresponding block */
func NewProofOfWork (b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256 - targetBits))
	pow := &ProofOfWork{b, target}
	return pow
}

/* Create a new Coinbase transaction */
/* Coinbase transaction is a special type of transaction, */
/* which doesn't require any previous input for validation */
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		/* Miner gets the reward for creating blocks */
		data = fmt.Sprintf("Reward to '%s'", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.setID()

	return &tx
}
