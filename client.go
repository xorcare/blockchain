package blockchain

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Go client blockchain.info api

const (
	apiRootTor = "https://blockchainbdgpzk.onion"
	apiRootNet = "https://blockchain.info"
)

type Client struct {
	*http.Client
	ApiRoot string
}

func (c *Client) doRequest(path string, i interface{}) (e error) {
	fullPath := c.ApiRoot + path

	response, e := c.Get(fullPath)
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
	return &Client{Client: &http.Client{}, ApiRoot: apiRootNet}
}

func NewTor() *Client {
	return &Client{Client: &http.Client{}, ApiRoot: apiRootTor}
}
