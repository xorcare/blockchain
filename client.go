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
	"net/url"
)

const (
	// The service address in the tor network
	ApiRootTor = "https://blockchainbdgpzk.onion"

	// The service address in the network
	ApiRootNet = "https://blockchain.info"
)

// The basic structure of the client
type Client struct {
	Http    *http.Client
	ApiRoot string
}

// Method to query the API. Automatically attempts to convert the response into the supplied type.
func (c *Client) DoRequest(path string, i interface{}, params map[string]string) (e error) {
	params["format"] = "json"
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Set(k, v)
	}

	fullPath := c.ApiRoot + path + "?" + (urlValues.Encode())
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

// Creating a new client running in the Internet
func New() *Client {
	return &Client{Http: &http.Client{}, ApiRoot: ApiRootNet}
}

// Create a new customer operating in the tor network
func NewTor() *Client {
	return &Client{Http: &http.Client{}, ApiRoot: ApiRootTor}
}
