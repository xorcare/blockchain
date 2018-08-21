// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain // import "github.com/xorcare/blockchain"

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Version api client.
	Version = "1.0"

	//Package package name const
	Package = "blockchain"

	// UserAgent is the header string used to identify this package.
	UserAgent = Package + "-api-go-client/" + Version + " (go; github; +https://git.io/fNhcB)"

	// BasePath the root address in the network.
	BasePath = "https://blockchain.info"

	// BasePathTor the root address in the tor network.
	BasePathTor = "https://blockchainbdgpzk.onion"
)

// Errors it is a set of errors returned when working with the package.
var (
	ErrAIW = errors.New(Package + ": address is wrong")
	ErrBEW = errors.New(Package + ": block height is wrong")
	ErrBHW = errors.New(Package + ": block hash is wrong")
	ErrCGD = errors.New(Package + ": cannot get data on url")
	ErrCRR = errors.New(Package + ": could not read answer response")
	ErrIRS = errors.New(Package + ": incorrect response status")
	ErrNAP = errors.New(Package + ": no address(es) provided")
	ErrRPE = errors.New(Package + ": response parsing error")
	ErrTHW = errors.New(Package + ": transaction hash is wrong")
)

// MaxAddressesCount the maximum number of addresses that can be checked at a time using
// the address checking function or the balance.
var MaxAddressesCount = 100

// Options map contains the default settings for requests to the api.
var Options = map[string]string{"format": "json"}

func init() {
	MaxAddressesCount = len(addressesForTestings)
}

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	client *http.Client

	APIKey    string // API access key.
	BasePath  string // API endpoint base URL.
	UserAgent string // Optional additional User-Agent fragment.
}

func (c *Client) userAgent() string {
	c.UserAgent = strings.TrimSpace(c.UserAgent)
	if c.UserAgent == "" {
		return UserAgent
	}

	return UserAgent + " " + c.UserAgent
}

// Do to send an client request, which is then converted to the passed type.
func (c *Client) Do(path string, i interface{}, options map[string]string) error {
	options = ApproveOptions(options)
	options["format"] = "json"
	options["api_code"] = c.APIKey
	if options["api_code"] == "" {
		delete(options, "api_code")
	}

	values := url.Values{}
	for k, v := range options {
		values.Set(k, v)
	}

	req, e := http.NewRequest("GET", c.BasePath+path+"?"+(values.Encode()), nil)
	if e != nil {
		return c.err2(ErrCGD, e)
	}
	req.Header.Set("User-Agent", c.userAgent())

	resp, e := c.client.Do(req)
	if e != nil {
		return c.err3(ErrCGD, e, resp)
	}
	defer resp.Body.Close()

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return c.err3(ErrCRR, e, resp)
	}

	if resp.Status[0] != '2' {
		return c.err3(ErrIRS, errors.New(string(bytes)), resp)
	}

	if e = json.Unmarshal(bytes, &i); e != nil {
		return c.err3(ErrRPE, e, resp)
	}

	return nil
}

// ApproveOptions automatic check and substitution of options.
func ApproveOptions(options map[string]string) map[string]string {
	if options == nil {
		return Options
	}
	return options
}

// New creates a new client instance the network internet.
func New() *Client {
	return &Client{client: &http.Client{}, BasePath: BasePath}
}

// NewTor creates a new client instance the tor network.
func NewTor() *Client {
	return &Client{client: &http.Client{}, BasePath: BasePathTor}
}

// SetClient http client setter.
func (c *Client) SetClient(client *http.Client) {
	if client == nil {
		panic(Package + ": impossible install the client as nil")
	}
	c.client = client
}
