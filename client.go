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
)

const (
	// APIRootTor the root Address in the tor network
	APIRootTor = "https://blockchainbdgpzk.onion"

	// APIRootNet the root Address in the network
	APIRootNet = "https://blockchain.info"
)

var (
	RRE = errors.New("could not read answer response")
	RSE = errors.New("incorrect response status")
	RPE = errors.New("response parsing error")
	WAE = errors.New("address is wrong")
	PAE = errors.New("no address(es) provided")
	CDE = errors.New("cannot get data on url")
	THW = errors.New("transaction hash is wrong")
	BEW = errors.New("block height is wrong")
	BHW = errors.New("block hash is wrong")
)

// Client specifies the mechanism by which individual API requests are made.
type Client struct {
	http    *http.Client
	apiRoot string
}

type Error struct {
	// ErrorMain error information from the standard package error set,
	ErrorMain error
	// ErrorExec information about the error that occurred during
	// the operation of the standard library or external packages
	ErrorExec error
	Response  *http.Response
	Address   *string
}

func (e Error) Error() string {
	return e.ErrorMain.Error()
}

func setErrorOne(e error) *Error {
	return setError(e, nil, nil, nil)
}

func setError(errorMain error, errorExec error, response *http.Response, address *string) *Error {
	if errorMain == nil {
		return nil
	}

	return &Error{
		ErrorMain: errorMain,
		ErrorExec: errorExec,
		Response:  response,
		Address:   address,
	}
}

// DoRequest to send an http request, which is then converted to the passed type.
func (c *Client) DoRequest(path string, i interface{}, params map[string]string) *Error {
	params["format"] = "json"
	urlValues := url.Values{}
	for k, v := range params {
		urlValues.Set(k, v)
	}

	uri := c.apiRoot + path + "?" + (urlValues.Encode())
	response, e := c.http.Get(uri)
	if e != nil {
		return setError(CDE, e, response, nil)
	}

	defer response.Body.Close()

	bytes, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return setError(RRE, e, response, nil)
	}

	if response.Status[0] != '2' {
		return setError(RSE, e, response, nil)
	}

	e = json.Unmarshal(bytes, &i)
	if e != nil {
		return setError(RPE, e, response, nil)
	}

	return nil
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
