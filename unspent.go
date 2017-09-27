// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"errors"
	"strings"
)

// UnspentOutputs the set of unspent outputs
type UnspentOutputs struct {
	Notice         string           `json:"notice,omitempty"`
	UnspentOutputs []*UnspentOutput `json:"unspent_outputs"`
}

// UnspentOutput the basic structure unspent outputs
type UnspentOutput struct {
	TxAge     string `json:"tx_age"`
	TxHash    string `json:"tx_hash"`
	TxIndex   uint64 `json:"tx_index"`
	TxOutputN uint64 `json:"tx_output_n"`
	Script    string `json:"script"`
	Value     int64  `json:"value"`
}

// GetUnspent specifies the mechanism by getting unspent outputs multiple addresses
func (c *Client) GetUnspent(addresses []string, params ...map[string]string) (response *UnspentOutputs, e error) {
	if len(addresses) == 0 {
		return nil, errors.New("No Address Provided")
	}

	options := map[string]string{"active": strings.Join(addresses, "|")}
	if len(params) > 0 {
		for k, v := range params[0] {
			options[k] = v
		}
	}
	response = &UnspentOutputs{}
	e = c.DoRequest("/unspent", response, options)

	return
}
