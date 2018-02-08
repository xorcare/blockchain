// Copyright 2017 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestGetAddress(t *testing.T) {
	response, e := New().GetAddress("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
	if e != nil {
		t.Fatal(e)
	}

	if response.Address != "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" {
		t.Fatal("Failed check address in the response")
	}

	if response.Hash160 != "62e907b15cbf27d5425399ebf6f0fb50ebb88f18" {
		t.Fatal("Failed check Hash160 in the response")
	}

	if response.NTx < 1000 {
		t.Fatal("Failed check number of transactions")
	}

	if len(response.Txs) < 50 {
		t.Fatal("Failed check count of transactions")
	}
}

func TestGetAddresses(t *testing.T) {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	}

	response, e := New().GetAddresses(addresses)
	if e != nil {
		t.Fatal(e)
	}

	if len(response.Txs) < 50 {
		t.Fatal("Failed check Txs")
	}

	for i := range response.Addresses {
		addr := response.Addresses[i]

		switch addr.Address {
		case "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa":
			if addr.NTx < 1105 {
				t.Fatal("Failed check number of transactions")
			}
		case "12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX":
			if addr.NTx < 57 {
				t.Fatal("Failed check number of transactions")
			}
		default:
			t.Fatal("Do not ordered address: " + addr.Address)
		}

		if len(addr.Txs) != 0 {
			t.Fatal("Failed check count of transactions")
		}
	}
}

func TestGetAddressesOneAddress(t *testing.T) {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	}

	response, e := New().GetAddresses(addresses)
	if e != nil {
		t.Fatal(e)
	}

	addr := response.Addresses[0]
	if addr.Address != "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" {
		t.Fatal("Failed check address in the addr")
	}

	if addr.NTx < 1000 {
		t.Fatal("Failed check number of transactions")
	}

	if len(addr.Txs) != 0 {
		t.Fatal("Wrong count of transactions")
	}
}

func TestGetAddressMoreParams(t *testing.T) {
	response, e := New().GetAddress("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(e)
	}

	if len(response.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestGetAddressesMoreParams(t *testing.T) {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	}

	response, e := New().GetAddresses(addresses, map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(e)
	}

	if len(response.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestAddressesBadParams(t *testing.T) {
	_, e := New().GetAddresses([]string{})
	if e == nil {
		t.Fatal("There must be a mistake")
	}

	_, e = New().GetAddress("")
	if e == nil {
		t.Fatal("There must be a mistake")
	}
}
