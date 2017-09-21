package blockchain

import (
	"strings"
)

type Balances map[string]*Balance

type Balance struct {
	FinalBalance  int `json:"final_balance"`
	NTx           int `json:"n_tx"`
	TotalReceived int `json:"total_received"`
}

func (c *Client) GetBalance(addresses []string) (response *Balances, e error) {
	var path = "/balance?active=" + strings.Join(addresses, "|")
	response = &Balances{}
	e = c.DoRequest(path, response)

	return
}
