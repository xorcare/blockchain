// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "strings"

// Balances structure of the response to the balance request.
type Balances map[string]Balance

// Balance describes the available data at the same address
// when you request a balance.
type Balance struct {
	FinalBalance  uint64 `json:"final_balance"`
	NTx           uint64 `json:"n_tx"`
	TotalReceived uint64 `json:"total_received"`
}

func (b Balances) append(balances Balances) {
	if balances != nil {
		for k, v := range balances {
			b[k] = v
		}
	}
}

// GetBalance the method for obtaining the balance of one or more
// addresses. For times check out the better not more than
// MaxAddressesCount locations.
func (c *Client) GetBalance(addresses []string) (resp Balances, e error) {
	if e = c.checkAddresses(addresses); e != nil {
		return
	}

	resp = Balances{}
	return resp, c.Do("/balance", &resp, map[string]string{
		"active": strings.Join(addresses, "|"),
	})
}

// GetBalanceImp the improved method for obtaining the balance of one or more addresses.
// Has no limit on the number of addresses.
func (c *Client) GetBalanceImp(addresses []string) (resp Balances, e error) {
	if e = c.checkAddresses(addresses); e != nil {
		return
	}

	resp = Balances{}
	if len(addresses) > MaxAddressesCount {
		count := MaxAddressesCount
		resp2, e := c.GetBalanceImp(addresses[:count])
		if e != nil {
			return resp, e
		}

		resp.append(resp2)
		resp3, e := c.GetBalanceImp(addresses[count:])
		if e != nil {
			return resp, e
		}
		resp.append(resp3)
	} else {
		return c.GetBalance(addresses)
	}

	return
}
