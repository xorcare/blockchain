package blockchain

import "strings"

type Address struct {
	Hash160       string `json:"hash160"` // Does not exist in the case multiaddr
	Address       string `json:"address"`
	NTx           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
	TotalSent     uint64 `json:"total_sent"`
	FinalBalance  uint64 `json:"final_balance"`
	Txs           []*Tx  `json:"txs"` // Does not exist in the case multiaddr

	// Exist in the case multiaddr
	ChangeIndex  uint64 `json:"change_index"`
	AccountIndex uint64 `json:"account_index"`
}

type MultiAddr struct {
	RecommendIncludeFee bool       `json:"recommend_include_fee"`
	SharedcoinEndpoint  string     `json:"sharedcoin_endpoint"`
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

func (c *Client) GetAddress(address string) (response *Address, e error) {
	var path = "/address/" + address + "?format=json"
	response = &Address{}
	e = c.DoRequest(path, response)

	return
}

func (c *Client) GetAddresses(addresses []string) (response *MultiAddr, e error) {
	var path = "/multiaddr?active=" + strings.Join(addresses, "|")
	response = &MultiAddr{}
	e = c.DoRequest(path, response)

	return
}
