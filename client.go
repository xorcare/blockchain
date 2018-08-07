// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// Version api client
	Version = "1.0"

	// UserAgent is the header string used to identify this package.
	UserAgent = "blockchain-api-v1-client-go/" + Version + " (go; github; +https://git.io/v5dN0)"

	// BasePath the root address in the network
	BasePath = "https://blockchain.info"

	// BasePathTor the root address in the tor network
	BasePathTor = "https://blockchainbdgpzk.onion"
)

// Errors it is a set of errors returned when working with the package
var (
	ErrAIW = errors.New("address is wrong")
	ErrNAP = errors.New("no address(es) provided")
	ErrBEW = errors.New("block height is wrong")
	ErrBHW = errors.New("block hash is wrong")
	ErrCGD = errors.New("cannot get data on url")
	ErrCRR = errors.New("could not read answer response")
	ErrIRS = errors.New("incorrect response status")
	ErrRPE = errors.New("response parsing error")
	ErrTHW = errors.New("transaction hash is wrong")
)

// Options map contains the default settings for requests to the api
var Options = map[string]string{"format": "json"}

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	client *http.Client
	error  *Error

	APIKey    string // API access key
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment
}

// GetLastError returns the data set of the last error that occurred
// representing an error condition, with the nil value representing no error.
func (c *Client) GetLastError() *Error {
	defer func(c *Client) {
		c.error = nil
	}(c)
	return c.error
}

func (c *Client) userAgent() string {
	if c.UserAgent == "" {
		return UserAgent
	}

	return UserAgent + " " + c.UserAgent
}

func (c *Client) setErrorOne(errorMain error) error {
	return c.setError(errorMain, nil, nil, nil)
}

func (c *Client) setErrorTwo(errorMain error, errorExec error) error {
	return c.setError(errorMain, errorExec, nil, nil)
}

func (c *Client) setErrorThree(errorMain error, errorExec error, response *http.Response) error {
	return c.setError(errorMain, errorExec, response, nil)
}

func (c *Client) setError(errorMain error, errorExec error, response *http.Response, address *string) error {
	c.error = nil

	if errorMain == nil {
		return nil
	}

	c.error = &Error{
		ErrMain:  errorMain,
		ErrExec:  errorExec,
		Response: response,
		Address:  address,
	}

	return errorMain
}

// Do to send an client request, which is then converted to the passed type.
func (c *Client) Do(path string, i interface{}, options map[string]string) error {
	options = ApproveOptions(options)
	options["format"] = "json"
	options["api_code"] = c.APIKey
	values := url.Values{}
	for k, v := range options {
		values.Set(k, v)
	}

	req, e := http.NewRequest("GET", c.BasePath+path+"?"+(values.Encode()), nil)
	if e != nil {
		return c.setErrorTwo(ErrCGD, e)
	}
	req.Header.Set("User-Agent", c.userAgent())

	resp, e := c.client.Do(req)
	if e != nil {
		return c.setErrorThree(ErrCGD, e, resp)
	}
	defer resp.Body.Close()

	bytes, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		return c.setErrorThree(ErrCRR, e, resp)
	}

	if resp.Status[0] != '2' {
		return c.setErrorThree(ErrIRS, e, resp)
	}

	if e = json.Unmarshal(bytes, &i); e != nil {
		return c.setErrorThree(ErrRPE, e, resp)
	}

	return nil
}

// Error data structure describing the error
type Error struct {
	// Address wrong address
	Address *string
	// ErrMain error information from the standard package error set,
	ErrMain error
	// ErrExec information about the error that occurred during
	// the operation of the standard library or external packages
	ErrExec error
	// Response http response
	Response *http.Response
}

// Error compatibility with error interface
func (e Error) Error() string {
	return e.ErrMain.Error()
}

// ApproveOptions automatic check and substitution of options
func ApproveOptions(options map[string]string) map[string]string {
	if options == nil {
		return Options
	}
	return options
}

// New specifies the mechanism by create newClient client the network internet
func New() *Client {
	return &Client{client: &http.Client{}, BasePath: BasePath}
}

// NewTor specifies the mechanism by create newClient client the network internet
func NewTor() *Client {
	return &Client{client: &http.Client{}, BasePath: BasePathTor}
}

// SetHTTP http client setter
func (c *Client) SetHTTP(cli *http.Client) {
	c.client = cli
}
