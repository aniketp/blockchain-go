## Blockchain-implementation

An implementation of blockchain in Golang.

## Demo
Initial (now redundant) approach was to store the Blocks containing SHA-256
Hashes of the current and previous blocks along with the data (usually the
transactions).

Later revisions added a CLI support using Golang's builtin `os.Args` package.

<p float="left">
  <img src="images/initial.gif" width="440" height="400" />
  <img src="images/second.gif" width="440" height="400" />
</p>

## Blockchain-CLI

The functionality to create custom transactions and store them as a ledger was
introduced after [66b9b8d](https://github.com/aniketp/blockchain/tree/66b9b8d9d4728b12bd4bb8a940f9fd6ea6485163).

### Dependencies
```
 $ go get github.com/boltdb/bolt
 $ go get github.com/eiannone/keyboard
```

### Installation & Usage
``` bash
 ▶ go build

 Usage: blockchain [addblock] [printchain ...]
```

### Hashing algorithm

Current restriction on calculating block hash is to have a certain length of 0s (18 as of now, see [here](https://www.blockchain.com/btc/block/000000000000000000e907ebdb890c7f46c0649829b60e98ff5cb5e2b83fcc77)) at the beginning. However, it is technically impossible
to calculate reasonable amount of hashes in a minute (generally takes years). Keeping
that in mind, I've set the limit to 4.

The hash function repeatedly calculates the hashes in order until it finds one
which satisfies the restriction.

``` Go
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
```
The `nonce` is incremented after each iteration to inculcate the randomness in
hash generation.

### Block structure

The `block` structure is defined as follows.
``` Go
type Block struct {
       Timestamp  	int64		/* Time of block creation */
       Data		[]byte		/* Valuable info in the block */
       PrevBlockHash	[]byte		/* Hash of the previous block */
       Hash 		[]byte		/* Hash of the block */
       Nonce		int		/* Random nonce for PoW */
}
```

### Proof of Work
POW is important in a sense that it enables us to validate whether the existing
transactions are frivolous or are they actually correct. The Proof of Work algorithm
validates whether the individual instances of every block satisfies the length
restriction and that the sha256 of its data and timestamp is matches the pre-computed value.

### Implementations TODO
* Extract unspent coins
* Network interaction
* Bitcoin client nodes
* Private wallets

### Credits
Thanks to the wonderful Golang documentation obviously! Apart from it, ideas were
picked up from these sources.

* [The Bitcoin Paper](https://bitcoin.org/bitcoin.pdf)
* [Blockchain.com](https://bitcoin.org/bitcoin.pdf)
* [Medium/Proof of Stake blockchain ](https://medium.com/@mycoralhealth/code-your-own-proof-of-stake-blockchain-in-go-610cd99aa658)
* [Medium/Code your own blockchain](https://medium.com/@mycoralhealth/code-your-own-blockchain-in-less-than-200-lines-of-go-e296282bcffc)

### Feeling Naughty?
Set `targetBits` in [utils.go](utils.go) as something above 20 (say 25) and watch your CPU blow up :stuck_out_tongue:
