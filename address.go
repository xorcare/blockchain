// Copyright 2017 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"strings"
)

// Address description of the address structure returned from the API,
// Some fields in some cases may be empty or absent.
type Address struct {
	// Exist only in the case address
	Hash160 string `json:"hash160,omitempty"`

	Address       string `json:"address"`
	NTx           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  uint64 `json:"final_balance"`
	Txs           []Tx   `json:"txs,omitempty"`

	// Exist only in the case multiaddr
	ChangeIndex  uint64 `json:"change_index,omitempty"`
	AccountIndex uint64 `json:"account_index,omitempty"`
}

// MultiAddr structure of the result when querying multiple addresses
type MultiAddr struct {
	RecommendIncludeFee bool      `json:"recommend_include_fee,omitempty"`
	SharedcoinEndpoint  string    `json:"sharedcoin_endpoint,omitempty"`
	Wallet              Wallet    `json:"wallet"`
	Addresses           []Address `json:"addresses"`
	Txs                 []Tx      `json:"txs"`
	Info                Info      `json:"info"`
}

// Wallet summary data about the requested addresses
type Wallet struct {
	NTx           uint64 `json:"n_tx"`
	NTxFiltered   uint64 `json:"n_tx_filtered"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  uint64 `json:"final_balance"`
}

// SymbolLocal ...
type SymbolLocal struct {
	Code               string  `json:"code"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Conversion         float64 `json:"conversion"`
	SymbolAppearsAfter bool    `json:"symbolAppearsAfter"`
	Local              bool    `json:"local"`
}

// SymbolBtc ...
type SymbolBtc struct {
	Code               string  `json:"code"`
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Conversion         float64 `json:"conversion"`
	SymbolAppearsAfter bool    `json:"symbolAppearsAfter"`
	Local              bool    `json:"local"`
}

// Info ...
type Info struct {
	NConnected  uint64      `json:"nconnected"`
	Conversion  float64     `json:"conversion"`
	SymbolLocal SymbolLocal `json:"symbol_local"`
	SymbolBtc   SymbolBtc   `json:"symbol_btc"`
	LatestBlock LatestBlock `json:"latest_block"`
}

// GetAddress alias GetAddressAdv without additional parameters
func (c *Client) GetAddress(address string) (*Address, error) {
	return c.GetAddressAdv(address, map[string]string{})
}

// GetAddressAdv is a mechanism which is used to obtain information about the address
func (c *Client) GetAddressAdv(address string, params ...map[string]string) (response *Address, e error) {
	addressLength := len(address)
	if address == "" || addressLength > 35 || addressLength < 26 {
		return nil, c.setError(WAE, nil, nil, nil)
	}

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

// GetAddresses alias GetAddressesAdv without additional parameters
func (c *Client) GetAddresses(addresses []string) (*MultiAddr, error) {
	return c.GetAddressesAdv(addresses, map[string]string{})
}

func (c *Client) CheckAddresses(addresses []string) (e error) {
	if len(addresses) == 0 {
		return c.setErrorOne(PAE)
	}

	for _, addr := range addresses {
		addressLength := len(addr)
		if addr == "" || addressLength > 35 || addressLength < 26 {
			return c.setError(WAE, nil, nil, &addr)
		}
	}

	return
}

// GetAddressesAdv is a mechanism which is used to obtain information about the addresses
func (c *Client) GetAddressesAdv(addresses []string, params ...map[string]string) (multiAddr *MultiAddr, e error) {
	e = c.CheckAddresses(addresses)
	if e != nil {
		return
	}

	options := map[string]string{"active": strings.Join(addresses, "|")}
	if len(params) > 0 {
		for k, v := range params[0] {
			options[k] = v
		}
	}

	multiAddr = &MultiAddr{}
	e = c.DoRequest("/multiaddr", multiAddr, options)

	return
}
