go
package main

import (
	"time"
	"crypto/sha256"
	"strconv"
	"log"
)

// Block represents block in blockchain
type Block struct {
	timestamp     int64
	data          []byte
	prevBlockHash []byte
	hash          []byte
}

// calculates and sets block hash
func (b *Block) SetHash() {
	
	h := sha256.New()
	timestamp := []byte(strconv.FormatInt(b.timestamp, 10))
	var data []byte
	
	data = append(data, b.prevBlockHash...)
	data = append(data, b.data...)
	data = append(data, timestamp...)

	_, err := h.Write(data)
	if err != nil {
		log.Panic(err)
	}

	b.hash = h.Sum(nil)
}

// creates and returns block based on data and previous block
func NewBlock(data string, prevBlockHash []byte) *Block {

	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte("")}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}