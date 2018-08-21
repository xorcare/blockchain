// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestClient_CheckAddress(t *testing.T) {
	t.Parallel()
	c := New()
	if e := c.checkAddress(""); e.Error() != ErrAIW.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestClient_CheckAddresses(t *testing.T) {
	t.Parallel()
	c := New()
	if e := c.checkAddresses([]string{}); e.Error() != ErrNAP.Error() {
		t.Fatal("incorrect error: " + e.Error())
	}
}

func TestValidateBitcoinAddress(t *testing.T) {
	t.Parallel()
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
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinAddress(test.address) != test.result {
				t.Fatalf("validate test failed address: %s", test.address)
			}
		})
	}
}

func TestValidateBitcoinXpub(t *testing.T) {
	t.Parallel()
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
	for _, test := range tests {
		t.Run(test.address, func(t *testing.T) {
			if ValidateBitcoinXpub(test.address) != test.result {
				t.Fatalf("validate test failed xpub: %s", test.address)
			}
		})
	}
}
