// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import (
	"encoding/json"
	"testing"
)

func TestClient_GetAddress(t *testing.T) {
	resp, e := newClient().GetAddress("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
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

	tests := []struct {
		address string
	}{
		// Unique addresses
		{"1111111111111111111114oLvT2"},
		// one signature addresses
		{"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9"},
		{"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy"},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd"},
		// multi signature addresses
		{"3BGvENmRvUNaiEFPeUdNQgqebznN7Vqeqk"},
		{"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa"},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq"},
	}
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			resp, e := newClient().GetAddress(test.address)
			if e != nil {
				t.Fatal(e)
			}

			if resp == nil || resp.Address != test.address {
				t.Fatalf("GetAddress test failed: %s", test.address)
			}
		})
	}
}

func TestClient_GetAddresses(t *testing.T) {
	t.Log("Max addresses count:", GetMaxAddressesCount())

	c := newClient()
	resp, e := c.GetAddresses(addressesForTestings)
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

	addresses := []string{
		// Unique addresses
		"1111111111111111111114oLvT2",
		// one signature addresses
		"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9",
		"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy",
		"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd",
		// multi signature addresses
		"3BGvENmRvUNaiEFPeUdNQgqebznN7Vqeqk",
		"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa",
		"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq",
		// xpub addresses
		"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz",
		"xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt",
		"xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon",
	}

	resp, e = c.GetAddresses(addresses)
	if e != nil || c.GetLastError() != nil {
		t.Fatal(e)
	}
}

func TestGetAddressesOneAddress(t *testing.T) {
	addresses := []string{
		"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
	}

	resp, e := newClient().GetAddresses(addresses)
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
	resp, e := newClient().GetAddressAdv("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa", map[string]string{"offset": "2147483647"})
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

	resp, e := newClient().GetAddressesAdv(addresses, map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(e)
	}

	if len(resp.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestAddressesBadOptions(t *testing.T) {
	if _, e := newClient().GetAddresses([]string{}); e == nil {
		t.Fatal("There must be a mistake")
	}

	if _, e := newClient().GetAddressAdv("", nil); e == nil {
		t.Fatal("There must be a mistake")
	}
}

func TestClient_CheckAddress(t *testing.T) {
	c := New()
	if e := c.checkAddress(""); e != ErrAIW {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckAddresses(t *testing.T) {
	c := New()
	if e := c.checkAddresses([]string{}); e != ErrNAP {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestValidateBitcoinAddress(t *testing.T) {
	tests := []struct {
		address string
		result  bool
	}{
		// Unique addresses
		{"1111111111111111111114oLvT2", true},
		// one signature addresses
		{"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", true},
		{"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy", true},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", true},
		// multi signature addresses
		{"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa", true},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", true},
		// bad addresses
		{"", false},
		{"1111111111111111111114oLvT", false},
		{"0111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
		{"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz", false},
		{"xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon", false},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

func TestValidateBitcoinXpub(t *testing.T) {
	tests := []struct {
		address string
		result  bool
	}{
		// good xpub
		{"xpub6CUGRUonZSQ4TWtTMmzXdrXDtypWKiKrhko4egpiMZbpiaQL2jkwSB1icqYh2cfDfVxdx4df189oLKnC5fSwqPfgyP3hooxujYzAu3fDVmz", true},
		{"xpub6FnCn6nSzZAw5Tw7cgR9bi15UV96gLZhjDstkXXxvCLsUXBGXPdSnLFbdpq8p9HmGsApME5hQTZ3emM2rnY5agb9rXpVGyy3bdW6EEgAtqt", true},
		{"xpub6DF8uhdarytz3FWdA8TvFSvvAh8dP3283MY7p2V4SeE2wyWmG5mg5EwVvmdMVCQcoNJxGoWaU9DCWh89LojfZ537wTfunKau47EL2dhHKon", true},
		// bad xpub
		{"", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
		{"1111111111111111111114oLvT2", false},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", false},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
	}
	t.Parallel()
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinXpub(test.address) != test.result {
				t.Fatalf("validate test failed xpub: %s", test.address)
			}
		})
	}
}

func BenchmarkAddressUnmarshal(b *testing.B) {
	b.StopTimer()
	resp, e := newClient().GetAddressAdv("16rCmCmbuWDhPjWTrpQGaU3EPdZF7MTdUk", map[string]string{})
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
