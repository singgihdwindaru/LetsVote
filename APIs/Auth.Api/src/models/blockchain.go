package models

import (
	"bytes"
	"crypto/sha256"
	"time"
)

type Block struct {
	Timestamp         int64
	PreviousBlockHash []byte
	CurrentBlockHash  []byte
	AllData           []byte
}

type blockchain struct {
	Blocks []*Block
}

func NewBlockchain(data string, prevBlockHash []byte) *blockchain {
	return &blockchain{[]*Block{
		newBlock(data, prevBlockHash)},
	}
}

func newBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), prevBlockHash, []byte{}, []byte(data)}
	block.SetHash()
	return block
}

func (block *Block) SetHash() {
	// timestamp := []byte(strconv.FormatInt(block.Timestamp, 10))
	headers := bytes.Join([][]byte{block.PreviousBlockHash, block.AllData}, []byte{})
	hash := sha256.Sum256(headers)
	block.CurrentBlockHash = hash[:]
}

// create the method that adds a new block to a blockchain
func (blockchain *blockchain) AddBlock(data string) {
	PreviousBlock := blockchain.Blocks[len(blockchain.Blocks)-1]
	newBlock := newBlock(data, PreviousBlock.CurrentBlockHash)
	blockchain.Blocks = append(blockchain.Blocks, newBlock)
}
