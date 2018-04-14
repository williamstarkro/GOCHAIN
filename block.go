package main

import (
	"time"
	"crypto/sha256"
	"strconv"
	"log"
)

// Block represents block in blockchain
type Block struct {
	Timestamp     int64
	Data          []byte
	PrevBlockHash []byte
	Hash          []byte
}

// calculates and sets block hash
func (b *Block) SetHash() {
	
	h := sha256.New()
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	var data []byte
	
	data = append(data, b.PrevBlockHash...)
	data = append(data, b.Data...)
	data = append(data, timestamp...)

	_, err := h.Write(data)
	if err != nil {
		log.Panic(err)
	}

	b.Hash = h.Sum(nil)
}

// creates and returns block based on data and previous block
func NewBlock(Data string, PrevBlockHash []byte) *Block {

	block := &Block{time.Now().Unix(), []byte(Data), PrevBlockHash, []byte("")}
	block.SetHash()
	return block
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte("0"))
}