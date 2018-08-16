// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

// Block the structure of one specific block.
type Block struct {
	Hash       string `json:"hash"`
	Ver        int64  `json:"ver"`
	PrevBlock  string `json:"prev_block"`
	MrklRoot   string `json:"mrkl_root"`
	Time       uint64 `json:"time"`
	Bits       uint64 `json:"bits"`
	Fee        int64  `json:"fee"`
	Nonce      uint64 `json:"nonce"`
	NTx        uint64 `json:"n_tx"`
	Size       uint64 `json:"size"`
	BlockIndex uint64 `json:"block_index"`
	MainChain  bool   `json:"main_chain"`
	Height     uint64 `json:"height"`
	Tx         []Tx   `json:"tx"`
}

// LatestBlock the structure of the last block in the chain.
type LatestBlock struct {
	Hash       string   `json:"hash"`
	Time       uint64   `json:"time"`
	BlockIndex uint64   `json:"block_index"`
	Height     uint64   `json:"height"`
	TxIndexes  []uint64 `json:"txIndexes"`
}

// Blocks the structure of the set of blocks.
type Blocks struct {
	Blocks []Block `json:"blocks"`
}

// GetBlock get the block by the hash.
func (c *Client) GetBlock(hash string) (resp *Block, e error) {
	if hash == "" || len(hash) != 64 {
		return nil, c.setErrorOne(ErrBHW)
	}

	resp = &Block{}
	return resp, c.Do("/rawblock/"+hash, resp, nil)
}

// GetBlockHeight get the block at height.
func (c *Client) GetBlockHeight(height string) (resp *Blocks, e error) {
	if height == "" {
		return nil, c.setErrorOne(ErrBEW)
	}

	resp = &Blocks{}
	return resp, c.Do("/block-height/"+height, resp, nil)
}

// GetBlocks getting blocks at a certain height.
func (c *Client) GetBlocks(height string) (resp *Blocks, e error) {
	if height == "" {
		return nil, c.setErrorOne(ErrBEW)
	}

	resp = &Blocks{}
	return resp, c.Do("/blocks/"+height, resp, nil)
}

// GetLatestBlock receive the latest block of the main chain.
func (c *Client) GetLatestBlock() (resp *LatestBlock, e error) {
	resp = &LatestBlock{}
	return resp, c.Do("/latestblock", resp, nil)
}
