// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

// Tx the basic structure of the transaction
type Tx struct {
	Result      int64    `json:"result"`
	Ver         int64    `json:"ver"`
	Size        uint64   `json:"size"`
	Inputs      []Inputs `json:"inputs"`
	Time        uint64   `json:"time"`
	BlockHeight uint64   `json:"block_height"`
	TxIndex     uint64   `json:"tx_index"`
	VinSz       uint64   `json:"vin_sz"`
	Hash        string   `json:"hash"`
	VoutSz      uint64   `json:"vout_sz"`
	RelayedBy   string   `json:"relayed_by"`
	Out         []Out    `json:"out"`
	Weight      uint64   `json:"weight"`
	Fee         int64    `json:"fee"`
	LockTime    int64    `json:"lock_time"`
	DoubleSpend bool     `json:"double_spend"`
	Balance     int64    `json:"balance"`
	Rbf         bool     `json:"rbf"`
}

// Inputs the main structure of the inputs
type Inputs struct {
	Sequence uint64  `json:"sequence"`
	Witness  string  `json:"witness"`
	PrevOut  PrevOut `json:"prev_out"`
	Script   string  `json:"script"`
}

// PrevOut ...
type PrevOut struct {
	AddrTagLink string `json:"addr_tag_link"`
	AddrTag     string `json:"addr_tag"`
	Spent       bool   `json:"spent"`
	TxIndex     uint64 `json:"tx_index"`
	Type        uint64 `json:"type"`
	Addr        string `json:"addr"`
	Value       uint64 `json:"value"`
	N           uint64 `json:"n"`
	Script      string `json:"script"`
}

// Out the main structure of the inputs
type Out struct {
	AddrTagLink string `json:"addr_tag_link"`
	AddrTag     string `json:"addr_tag"`
	Spent       bool   `json:"spent"`
	TxIndex     uint64 `json:"tx_index"`
	Type        uint64 `json:"type"`
	Addr        string `json:"addr"`
	Value       uint64 `json:"value"`
	N           uint64 `json:"n"`
	Script      string `json:"script"`
}

// Txs transaction set
type Txs struct {
	Txs []Tx `json:"txs"`
}

// GetTransaction get the transaction on its hash
func (c *Client) GetTransaction(transaction string) (response *Tx, e error) {
	if transaction == "" || len(transaction) != 64 {
		return nil, c.setErrorOne(THW)
	}

	response = &Tx{}
	e = c.DoRequest("/rawtx/"+transaction, response, nil)

	return
}

// GetUnconfirmedTransactions get the unconfirmed transactions
func (c *Client) GetUnconfirmedTransactions() (reaponse *Txs, e error) {
	reaponse = &Txs{}
	e = c.DoRequest("/unconfirmed-transactions", reaponse, nil)

	return
}
