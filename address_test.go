// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestClient_GetAddress(t *testing.T) {
	resp, e := newClient().GetAddress(firstBitcoinAddress)
	if e != nil {
		t.Fatal(customError(e))
	}

	if resp.Address != firstBitcoinAddress {
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
		// one signature addresses
		{"1111111111111111111114oLvT2"},
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
				t.Fatal(customError(e))
			}

			if resp == nil || resp.Address != test.address {
				t.Fatalf("GetAddress test failed: %s", test.address)
			}
		})
	}
}

func TestClient_GetAddresses(t *testing.T) {
	c := newClient()
	resp, e := c.GetAddresses(addressesForTestings)
	if e != nil {
		t.Fatal(customError(e))
	}

	if len(resp.Txs) < 50 {
		t.Fatal("Failed check Txs")
	}

	for _, addr := range resp.Addresses {
		switch addr.Address {
		case firstBitcoinAddress:
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
		// one signature addresses
		"1111111111111111111114oLvT2",
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
	if e != nil {
		t.Fatal(customError(e))
	} else {
		for _, v := range resp.Addresses {
			t.Logf("%34s %11d %11d %5d", v.Address, v.FinalBalance, v.TotalReceived, v.NTx)
		}
	}
}

func TestGetAddressesOneAddress(t *testing.T) {
	addresses := []string{
		firstBitcoinAddress,
	}

	resp, e := newClient().GetAddresses(addresses)
	if e != nil {
		t.Fatal(customError(e))
	}

	addr := resp.Addresses[0]
	if addr.Address != firstBitcoinAddress {
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
	resp, e := newClient().GetAddressAdv(firstBitcoinAddress, map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(customError(e))
	}

	if len(resp.Txs) != 0 {
		t.Fatal("Wrong count txs")
	}
}

func TestGetAddressesMoreOptions(t *testing.T) {
	addresses := []string{
		firstBitcoinAddress,
		"12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX",
	}

	resp, e := newClient().GetAddressesAdv(addresses, map[string]string{"offset": "2147483647"})
	if e != nil {
		t.Fatal(customError(e))
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
	if e := c.checkAddress(""); e.Error() != ErrAIW.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckAddresses(t *testing.T) {
	c := New()
	if e := c.checkAddresses([]string{}); e.Error() == ErrAIW.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestValidateBitcoinAddress(t *testing.T) {
	tests := []struct {
		address string
		result  bool
	}{
		// one signature addresses
		{"1111111111111111111114oLvT2", true},
		{"1dyoBoF5vDmPCxwSsUZbbYhA5qjAfBTx9", true},
		{"15yN7NPEpu82sHhB6TzCW5z5aXoamiKeGy", true},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", true},
		// multi signature addresses
		{"3B8SEgcT9JDVKUZvm8HoKX5Av3nnn7pHqa", true},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", true},
		// bad addresses
		{"", false},
		{"1111111111111111111114oLvT", false},
		{"1111111111111111111114iLvT", false},
		{"0111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
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
		{"1111111111111111111114oLvT2", false},
		{"xpub6DF8uhdarytz3FWdA8TvFSv", false},
		{"1F3EpcBBjVGaUuEJff9xYZBfBBALm1yfsd", false},
		{"3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
		{"xpub3KGPnzYshia2uSSz8BED2kSpx22bbGCkzq", false},
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
