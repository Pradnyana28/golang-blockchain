package blockchain

import (
	"bytes"
	"crypto/sha256"
)

// BlockChain ...
type BlockChain struct {
	Blocks []*Block
}

// Block ...
type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

// DeriveHash ...
func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

// CreateBlock ...
func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

// AddBlock ...
func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	new := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, new)
}

// Genesis ...
func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

// InitBlockChain ...
func InitBlockChain() *BlockChain {
	return &BlockChain{[]*Block{Genesis()}}
}
