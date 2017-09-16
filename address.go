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

	// Exist in the case multiaddr
	Weight      uint64 `json:"weight"`
	Fee         uint64 `json:"fee"`
	LockTime    uint64 `json:"lock_time"`
	DoubleSpend bool   `json:"double_spend"`
	Balance     uint64 `json:"balance"`
	Rbf         bool   `json:"rbf"`
}

type Inputs struct {
	Sequence uint64      `json:"sequence"`
	Witness  interface{} `json:"witness"`
	PrevOut  *PrevOut    `json:"prev_out"`

	// Exist in the case multiaddr
	Script string `json:"script"`
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

type LatestBlock struct {
	BlockIndex uint64 `json:"block_index"`
	Hash       string `json:"hash"`
	Height     uint64 `json:"height"`
	Time       uint64 `json:"time"`
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
	e = c.doRequest(path, response)

	return
}

func (c *Client) GetAddresses(addresses []string) (response *MultiAddr, e error) {
	var path = "/multiaddr?active=" + strings.Join(addresses, "|")
	response = &MultiAddr{}
	e = c.doRequest(path, response)

	return
}
