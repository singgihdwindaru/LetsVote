package models

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp         int64  // the time when the block was created
	PreviousBlockHash []byte // the hash of the previous block
	MyBlockHash       []byte // the hash of the current block
	AllData           []byte // the data or transactions (body info)
}

// Prepare the Blockchain data structure :
type Blockchain struct {
	Blocks []*Block // remember a blockchain is a series of blocks
}

// Now let's create a method for generating a hash of the block
// We will just concatenate all the data and hash it to obtain the block hash
func (block *Block) SetHash() {
	// timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))                                  // get the time and convert it into a unique series of digits
	headers := bytes.Join([][]byte{block.PreviousBlockHash, block.AllData}, []byte{}) // concatenate all the block data
	hash := sha256.Sum256(headers)                                                    // hash the whole thing
	block.MyBlockHash = hash[:]                                                       // now set the hash of the block
}

// Create a function for new block generation and return that block
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)} // the block is received
	block.SetHash()                                                           // the block is hashed
	return block                                                              // the block is returned with all the information in it
}

// /* let's now create the genesis block function that will return the first block. The genesis block is the first block on the chain */
// func NewGenesisBlock(value string) *Block {
// 	return NewBlock(value, []byte{}) // the genesis block is made with some data in it
// }

// create the method that adds a new block to a blockchain
func (blockchain *Blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1] // the previous block is needed, so let's get it
	newBlock := NewBlock(data, PreviousBlock.MyBlockHash)        // create a new block containing the data and the hash of the previous block
	blockchain.Blocks = append(blockchain.Blocks, newBlock)      // add that block to the chain to create a chain of blocks
}

/* Create the function that returns the whole blockchain and add the genesis to it first. the genesis block is the first ever mined block, so let's create a function that will return it since it does not exist yet */
func NewBlockchain(data string, prevBlockHash []byte) *Blockchain { // the function is created
	return &Blockchain{[]*Block{NewBlock(data, prevBlockHash)}} // the block is added first to the chain
}
