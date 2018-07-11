package main

type Blockchain struct {
	blocks		[]*Block
}

/* Add a new block to existing blockchain */
func (bc *Blockchain) Addblock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}
