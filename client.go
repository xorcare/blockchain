// Copyright 2017 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Golang client blockchain api -> https://blockchain.info/api

package blockchain

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

const (
	// APIRootTor the root address in the tor network
	APIRootTor = "https://blockchainbdgpzk.onion"

	// APIRootNet the root address in the network
	APIRootNet = "https://blockchain.info"
)

var (
	RRE = errors.New("could not read answer response")
	RSE = errors.New("incorrect response status")
	RPE = errors.New("response parsing error")
	WAE = errors.New("address is wrong")
	PAE = errors.New("no address(es) provided")
)

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	http      *http.Client
	apiRoot   string
	lastError *Error
}

type Error struct {
	error    error
	response *http.Response
	address  string
}

func (e Error) Error() string {
	return e.error.Error()
}

func (e Error) Response() http.Response {
	return *e.response
}
func (e Error) Address() string {
	return e.address
}

func (c *Client) CleanError() {
	c.lastError = nil
}

func (c *Client) GetLastError() (e *Error) {
	e = c.lastError
	c.CleanError()
	return
}

func (c *Client) setErrorResponse(e error, r *http.Response) *Error {
	c.lastError = &Error{
		error:    e,
		response: r,
	}

	return c.lastError
}
func (c *Client) setErrorAddress(e error, a string) *Error {
	c.lastError = &Error{
		error:   e,
		address: a,
	}
	return c.lastError
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
		c.setErrorResponse(e, response)
		return RRE
	}

	if response.Status[0] != '2' {
		c.setErrorResponse(RSE, response)
		return RSE
	}

	e = json.Unmarshal(bytes, &i)

	b := *c.setErrorResponse(RSE, response)

	if b == RSE {
		os.Exit(0)
	}

	if e == nil {
		c.CleanError()
	} else {
		c.setErrorResponse(RPE, response)
		return RPE
	}

	return
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
