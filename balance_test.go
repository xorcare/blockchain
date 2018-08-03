// Copyright 2017-2018 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestClient_GetBalance(t *testing.T) {
	t.Log("Max addresses count:", GetMaxAddressesCount())

	balances, e := newClient().GetBalance(addressesForTestings)
	if e != nil {
		t.Fatal(e)
	}

	for k, v := range balances {
		if v.FinalBalance > v.TotalReceived {
			t.Fatal("Error parsing total received")
		}

		if (v.FinalBalance > 0 || v.TotalReceived > 0) && v.NTx < 1 {
			t.Fatal("Received incorrect information about the balance or number of transactions")
		}

		t.Logf("%34s %11d %11d %5d", k, v.FinalBalance, v.TotalReceived, v.NTx)
	}

	t.Logf("Total addreses: %3d, total balances %3d", len(addressesForTestings), len(balances))
	if len(addressesForTestings) != len(balances) {
		t.Fatal("The number of addresses does not match the number received balances")
	}

	if _, e = newClient().GetBalance([]string{}); e == nil {
		t.Fatal("There must be a mistake")
	}
}
