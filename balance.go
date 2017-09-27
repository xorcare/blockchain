// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"errors"
	"strings"
)

// Structure of the result of the request balances
type Balances map[string]*Balance

// Structure Balance describes the available data at the same address when you request a balance
type Balance struct {
	FinalBalance  uint64 `json:"final_balance"`
	NTx           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
}

// The method for obtaining the balance of one or more addresses.
// For times check out the better not more than 200 locations.
func (c *Client) GetBalance(addresses []string) (response *Balances, e error) {
	if len(addresses) == 0 {
		return nil, errors.New("No Address Provided")
	}

	response = &Balances{}
	e = c.DoRequest("/balance", response, map[string]string{"active": strings.Join(addresses, "|")})

	return
}
