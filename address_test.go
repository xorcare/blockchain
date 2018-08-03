// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"encoding/json"
	"testing"
)

func TestClient_GetAddress(t *testing.T) {
	resp, e := New().GetAddress("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
	if e != nil {
		t.Fatal(e)
	}

	if resp.Address != "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" {
		t.Fatal("Failed check address in the response")
	}

	if resp.Hash160 != "62e907b15cbf27d5425399ebf6f0fb50ebb88f18" {
		t.Fatal("Failed check Hash160 in the response")
	}

	if resp.NTx < 1000 {
		t.Fatal("Failed check number of transactions")
	}

	if len(resp.Txs) < 50 {
		t.Fatal("Failed check count of transactions")
	}
}

func TestClient_GetAddresses(t *testing.T) {
	t.Log("Max addresses count:", GetMaxAddressesCount())

	resp, e := New().GetAddresses(addressesForTestings)
	if e != nil {
		t.Fatal(e)
	}

	if len(resp.Txs) < 50 {
		t.Fatal("Failed check Txs")
	}

	for _, addr := range resp.Addresses {
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
			if addr.NTx < 1 {
				t.Fatal("Failed check number of transactions")
			}
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

	resp, e := New().GetAddresses(addresses)
	if e != nil {
		t.Fatal(e)
	}

	addr := resp.Addresses[0]
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

func TestGetAddressMoreOptions(t *testing.T) {
	resp, e := New().GetAddressAdv("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(e)
	}

	if len(resp.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestGetAddressesMoreOptions(t *testing.T) {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
		"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	}

	resp, e := New().GetAddressesAdv(addresses, map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(e)
	}

	if len(resp.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestAddressesBadOptions(t *testing.T) {
	if _, e := New().GetAddresses([]string{}); e == nil {
		t.Fatal("There must be a mistake")
	}

	if _, e := New().GetAddressAdv("", nil); e == nil {
		t.Fatal("There must be a mistake")
	}
}

func TestClient_GetAddress__Errors(t *testing.T) {
	c := New()
	if _, e := c.GetAddress(""); e != AIW {
		t.Fatal("incorrect error: " + e.Error())
	}

	if _, e := c.GetAddresses([]string{}); e != ANP {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckAddresses(t *testing.T) {
	c := New()
	if e := c.CheckAddresses([]string{}); e != ANP {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func BenchmarkAddressUnmarshal(b *testing.B) {
	b.StopTimer()
	resp, e := New().GetAddressAdv("16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk", map[string]string{})
	if e != nil {
		b.Fatal(e)
	}
	bytes, e2 := json.Marshal(resp)
	if e2 != nil {
		b.Fatal(e2)
	}

	address := &Address{}
	b.ReportAllocs()
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		e := json.Unmarshal(bytes, address)
		if e != nil {
			b.Fatal(e)
		}
	}
}
