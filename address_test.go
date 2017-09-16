package blockchain

import (
	"testing"
)

func TestGetAddress(t *testing.T) {
	c := New()
	response, e := c.GetAddress("1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa")
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

	c := New()
	response, e := c.GetAddresses(addresses)
	if e != nil {
		t.Fatal(e)
	}

	if response.Info == nil {
		t.Fatal("Failed check Info")
	}

	if response.Wallet == nil {
		t.Fatal("Failed check Wallet")
	}

	if len(response.Txs) < 50 {
		t.Fatal("Failed check Txs")
	}

	for i := range response.Addresses {
		addr := response.Addresses[i]
		switch addr.Address {
		case "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa":
			t.Logf("Check address: %s", addr.Address)

			if addr.NTx < 1105 {
				t.Fatal("Failed check number of transactions")
			}

			if len(addr.Txs) != 0 {
				t.Fatal("Failed check count of transactions")
			}
		case "12c6DSiU4Rq3P4ZxziKxzrL5LmMBrzjrJX":
			t.Logf("Check address: %s", addr.Address)

			if addr.NTx < 57 {
				t.Fatal("Failed check number of transactions")
			}

			if len(addr.Txs) != 0 {
				t.Fatal("Failed check count of transactions")
			}
		}

	}
}
