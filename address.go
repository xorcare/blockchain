// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "strings"

type Address struct {
	Hash160       string `json:"hash160,omitempty"` // Exist only in the case address
	Address       string `json:"address"`
	NTx           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  uint64 `json:"final_balance"`
	Txs           []*Tx  `json:"txs,omitempty"`
	ChangeIndex   uint64 `json:"change_index,omitempty"`  // Exist only in the case multiaddr
	AccountIndex  uint64 `json:"account_index,omitempty"` // Exist only in the case multiaddr
}

type MultiAddr struct {
	RecommendIncludeFee bool       `json:"recommend_include_fee,omitempty"`
	SharedcoinEndpoint  string     `json:"sharedcoin_endpoint,omitempty"`
	Wallet              *Wallet    `json:"wallet"`
	Addresses           []*Address `json:"addresses"`
	Txs                 []*Tx      `json:"txs"`
	Info                *Info      `json:"info"`
}

type Wallet struct {
	NTx           uint64 `json:"n_tx"`
	NTxFiltered   uint64 `json:"n_tx_filtered"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  uint64 `json:"final_balance"`
}

type SymbolLocal struct {
	Code               string  `json:"code"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Conversion         float64 `json:"conversion"`
	SymbolAppearsAfter bool    `json:"symbolAppearsAfter"`
	Local              bool    `json:"local"`
}

type SymbolBtc struct {
	Code               string  `json:"code"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Conversion         float64 `json:"conversion"`
	SymbolAppearsAfter bool    `json:"symbolAppearsAfter"`
	Local              bool    `json:"local"`
}

type Info struct {
	NConnected  uint64       `json:"nconnected"`
	Conversion  float64      `json:"conversion"`
	SymbolLocal *SymbolLocal `json:"symbol_local"`
	SymbolBtc   *SymbolBtc   `json:"symbol_btc"`
	LatestBlock *LatestBlock `json:"latest_block"`
}

func (c *Client) GetAddress(address string, params ...map[string]string) (response *Address, e error) {
	options := map[string]string{"format": "json"}
	if len(params) > 0 {
		for k, v := range params[0] {
			options[k] = v
		}
	}
	response = &Address{}
	e = c.DoRequest("/address/"+address, response, options)

	return
}

func (c *Client) GetAddresses(addresses []string, params ...map[string]string) (response *MultiAddr, e error) {
	options := map[string]string{"active": strings.Join(addresses, "|")}
	if len(params) > 0 {
		for k, v := range params[0] {
			options[k] = v
		}
	}
	response = &MultiAddr{}
	e = c.DoRequest("/multiaddr", response, options)

	return
}
