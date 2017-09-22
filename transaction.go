// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

type Tx struct {
	Result      int64     `json:"result"`
	Ver         uint64    `json:"ver"`
	Size        uint64    `json:"size"`
	Inputs      []*Inputs `json:"inputs"`
	Time        uint64    `json:"time"`
	BlockHeight uint64    `json:"block_height"`
	TxIndex     uint64    `json:"tx_index"`
	VinSz       uint64    `json:"vin_sz"`
	Hash        string    `json:"hash"`
	VoutSz      uint64    `json:"vout_sz"`
	RelayedBy   string    `json:"relayed_by"`
	Out         []*Out    `json:"out"`
	Weight      uint64    `json:"weight"`
	Fee         uint64    `json:"fee"`
	LockTime    uint64    `json:"lock_time"`
	DoubleSpend bool      `json:"double_spend"`
	Balance     int64     `json:"balance"`
	Rbf         bool      `json:"rbf"`
}

type Inputs struct {
	Sequence uint64      `json:"sequence"`
	Witness  interface{} `json:"witness"`
	PrevOut  *PrevOut    `json:"prev_out"`
	Script   string      `json:"script"`
}

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

type Txs struct {
	Txs []*Tx `json:"txs"`
}

func (c *Client) GetTransaction(transaction string) (response *Tx, e error) {
	response = &Tx{}
	e = c.DoRequest("/rawtx/"+transaction, response, map[string]string{"format": "json"})

	return
}

func (c *Client) GetUnconfirmedTransactions() (reaponse *Txs, e error) {
	reaponse = &Txs{}
	e = c.DoRequest("/unconfirmed-transactions", reaponse, map[string]string{"format": "json"})

	return
}
