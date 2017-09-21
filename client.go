// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Golang client blockchain api -> https://blockchain.info/api

package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ApiRootTor = "https://blockchainbdgpzk.onion"
	ApiRootNet = "https://blockchain.info"
)

type Client struct {
	Http    *http.Client
	ApiRoot string
}

func (c *Client) DoRequest(path string, i interface{}) (e error) {
	fullPath := c.ApiRoot + path

	response, e := c.Http.Get(fullPath)
	if e != nil {
		return
	}

	defer response.Body.Close()

	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}
	if response.Status[0] != '2' {
		return fmt.Errorf("Response error status %3s: %s", response.Status, string(bytes))
	}

	return json.Unmarshal(bytes, &i)
}

func New() *Client {
	return &Client{Http: &http.Client{}, ApiRoot: ApiRootNet}
}

func NewTor() *Client {
	return &Client{Http: &http.Client{}, ApiRoot: ApiRootTor}
}
