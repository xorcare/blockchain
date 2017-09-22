// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

type Block struct {
	Hash       string `json:"hash"`
	Ver        uint64 `json:"ver"`
	PrevBlock  string `json:"prev_block"`
	MrklRoot   string `json:"mrkl_root"`
	Time       uint64 `json:"time"`
	Bits       uint64 `json:"bits"`
	Fee        uint64 `json:"fee"`
	Nonce      uint64 `json:"nonce"`
	NTx        uint64 `json:"n_tx"`
	Size       uint64 `json:"size"`
	BlockIndex uint64 `json:"block_index"`
	MainChain  bool   `json:"main_chain"`
	Height     uint64 `json:"height"`
	Tx         []*Tx  `json:"tx"`
}

type LatestBlock struct {
	Hash       string   `json:"hash"`
	Time       uint64   `json:"time"`
	BlockIndex uint64   `json:"block_index"`
	Height     uint64   `json:"height"`
	TxIndexes  []uint64 `json:"txIndexes"`
}

type Blocks struct {
	Blocks []*Block `json:"blocks"`
}

func (c *Client) GetBlock(block string) (response *Block, e error) {
	response = &Block{}
	e = c.DoRequest("/rawblock/"+block, response, map[string]string{"format": "json"})

	return
}

func (c *Client) GetBlockHeight(blockHeight string) (response *Blocks, e error) {
	response = &Blocks{}
	e = c.DoRequest("/block-height/"+blockHeight, response, map[string]string{"format": "json"})

	return
}

func (c *Client) GetBlocks(blockHeight string) (response *Blocks, e error) {
	response = &Blocks{}
	e = c.DoRequest("/blocks/"+blockHeight, response, map[string]string{"format": "json"})

	return
}

func (c *Client) GetLatestBlock() (response *LatestBlock, e error) {
	response = &LatestBlock{}
	e = c.DoRequest("/latestblock", response, map[string]string{"format": "json"})

	return
}
