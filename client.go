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
	// APIRootTor the root address in the tor network
	APIRootTor = "https://blockchainbdgpzk.onion"

	// APIRootNet the root address in the network
	APIRootNet = "https://blockchain.info"
)

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	http    *http.Client
	apiRoot string
}

// DoRequest to send an http request, which is then converted to the passed type.
func (c *Client) DoRequest(path string, i interface{}, params map[string]string) (e error) {
	params["format"] = "json"
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Set(k, v)
	}

	fullPath := c.apiRoot + path + "?" + (urlValues.Encode())
	response, e := c.http.Get(fullPath)
	if e != nil {
		return
	}

	defer response.Body.Close()

	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return
	}
	if response.Status[0] != '2' {
		return fmt.Errorf("Response error status %3s: %c", response.Status, string(bytes))
	}

	return json.Unmarshal(bytes, &i)
}

// New specifies the mechanism by create new client the network internet
func New() *Client {
	return &Client{http: &http.Client{}, apiRoot: APIRootNet}
}

// NewTor specifies the mechanism by create new client the network internet
func NewTor() *Client {
	return &Client{http: &http.Client{}, apiRoot: APIRootTor}
}

// SetHTTP http client setter
func (c *Client) SetHTTP(cli *http.Client) {
	c.http = cli
}
