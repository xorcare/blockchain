// Copyright 2017 Vasiliy Vasilyuk. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

import "testing"

func TestGetUnconfirmedTransactions(t *testing.T) {
	txs, e := New().GetUnconfirmedTransactions()
	if e != nil {
		t.Fatal(e)
	}

	for i := range txs.Txs {
		tx := txs.Txs[i]

		if len(tx.Hash) != 64 {
			t.Fatal("Wrong length value on field 'Hash'")
		}

		if tx.Ver < 1 {
			t.Fatal("Wrong value on field 'Ver'")
		}

		if tx.TxIndex < 1 {
			t.Fatal("Wrong value on field 'TxIndex'")
		}

		if tx.Time < 1 {
			t.Fatal("Wrong value on field 'Time'")
		}

		if len(tx.Inputs) < 1 {
			t.Fatal("Wrong count items on field 'Inputs'")
		}

		if len(tx.Out) < 1 {
			t.Fatal("Wrong count items on field 'Out'")
		}
	}
}

func TestGetTransaction(t *testing.T) {
	tx, e := New().GetTransaction("4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b")
	if e != nil {
		t.Fatal(e)
	}
	checkFirstTx(t, *tx)
}

func checkFirstTx(t *testing.T, tx Tx) {
	t.Logf("Started: checkFirstTx")

	if tx.Hash != "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b" {
		t.Fatal("Wrong value on field 'Tx[0].Hash'")
	}

	if tx.LockTime != 0 {
		t.Fatal("Wrong value on field 'Tx[0].LockTime'")
	}

	if tx.Ver != 1 {
		t.Fatal("Wrong value on field 'Tx[0].Ver'")
	}

	if tx.Size != 204 {
		t.Fatal("Wrong value on field 'Tx[0].Size'")
	}

	if tx.Weight != 816 {
		t.Fatal("Wrong value on field 'Tx[0].Weight'")
	}

	if tx.Time != 1231006505 {
		t.Fatal("Wrong value on field 'Tx[0].Time'")
	}

	if tx.TxIndex != 14849 {
		t.Fatal("Wrong value on field 'Tx[0].TxIndex'")
	}

	if tx.VinSz != 1 {
		t.Fatal("Wrong value on field 'Tx[0].VinSz'")
	}

	if tx.RelayedBy != "0.0.0.0" {
		t.Fatal("Wrong value on field 'Tx[0].RelayedBy'")
	}

	inputs0 := tx.Inputs[0]

	if inputs0.Sequence != 4294967295 {
		t.Fatal("Wrong value on field 'Tx[0].Inputs[0].Sequence'")
	}

	inputScript := "04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6"
	inputScript += "f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73"
	if inputs0.Script != inputScript {
		t.Fatal("Wrong value on field 'Tx[0].Inputs[0].Script'")
	}

	checkFirstTxOut(t, tx.Out[0])

}

func checkFirstTxOut(t *testing.T, out Out) {
	t.Logf("Started: checkFirstTxOut")

	if out.Spent != false {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Spent'")
	}

	if out.TxIndex != 14849 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].TxIndex'")
	}

	if out.Type != 0 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Type'")
	}

	if out.Addr != "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Addr'")
	}

	if out.Value != 5000000000 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Value'")
	}

	if out.N != 0 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].N'")
	}

	outScript := "4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb"
	outScript += "649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac"
	if out.Script != outScript {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].N'")
	}
}
