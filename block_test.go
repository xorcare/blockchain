package blockchain

import (
	"testing"
)

func TestGetBlock(t *testing.T) {
	c := New()
	block, e := c.GetBlock("000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f")
	if e != nil {
		t.Fatal(e)
	}
	checkFirstBlock(t, block)
}

func TestGetBlocks(t *testing.T) {
	c := New()
	blockHeight, e := c.GetBlockHeight("0")
	if e != nil {
		t.Fatal(e)
	}

	if len(blockHeight.Blocks) != 1 {
		t.Fatal("Wrong count items of field 'Blocks'")
	}

	checkFirstBlock(t, blockHeight.Blocks[0])
}

func TestGetLatestBlock(t *testing.T) {
	c := New()
	block, e := c.GetLatestBlock()
	if e != nil {
		t.Fatal(e)
	}

	if len(block.Hash) != 64 {
		t.Fatal("Wrong length value on field 'Hash'")
	}

	if block.Time < 1 {
		t.Fatal("Wrong value on field 'Time'")
	}

	if block.BlockIndex < 1 {
		t.Fatal("Wrong value on field 'BlockIndex'")
	}

	if block.Height < 1 {
		t.Fatal("Wrong value on field 'Height'")
	}

	if len(block.TxIndexes) < 1 {
		t.Fatal("Wrong count items on field 'TxIndexes'")
	}
}

// First block on bitcoin blockchain
//{
//   "hash":"000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f",
//   "ver":1,
//   "prev_block":"0000000000000000000000000000000000000000000000000000000000000000",
//   "mrkl_root":"4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b",
//   "time":1231006505,
//   "bits":486604799,
//   "fee":0,
//   "nonce":2083236893,
//   "n_tx":1,
//   "size":285,
//   "block_index":14849,
//   "main_chain":true,
//   "height":0,
//   "tx":[
//      {
//         "lock_time":0,
//         "ver":1,
//         "size":204,
//         "inputs":[
//            {
//               "sequence":4294967295,
//               "witness":null,
//               "script":"04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6
//                          f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73"
//            }
//         ],
//         "weight":816,
//         "time":1231006505,
//         "tx_index":14849,
//         "vin_sz":1,
//         "hash":"4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b",
//         "vout_sz":1,
//         "relayed_by":"0.0.0.0",
//         "out":[
//            {
//               "addr_tag_link":"https:\/\/en.bitcoin.it\/wiki\/Genesis_block",
//               "addr_tag":"Genesis of Bitcoin",
//               "spent":false,
//               "tx_index":14849,
//               "type":0,
//               "addr":"1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa",
//               "value":5000000000,
//               "n":0,
//               "script":"4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb
//                          649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac"
//            }
//         ]
//      }
//   ]
//}

func checkFirstBlock(t *testing.T, block *Block) {
	if block.Hash != "000000000019d6689c085ae165831e934ff763ae46a2a6c172b3f1b60a8ce26f" {
		t.Fatal("Wrong value on field 'Hash'")
	}

	if block.Ver != 1 {
		t.Fatal("Wrong value on field 'Ver'")
	}

	if block.PrevBlock != "0000000000000000000000000000000000000000000000000000000000000000" {
		t.Fatal("Wrong value on field 'PrevBlock'")
	}

	if block.Time != 1231006505 {
		t.Fatal("Wrong value on field 'PrevBlock'")
	}

	if block.NTx != 1 {
		t.Fatal("Wrong value on field 'NTx'")
	}

	if block.Size != 285 {
		t.Fatal("Wrong value on field 'Size'")
	}

	if block.BlockIndex != 14849 {
		t.Fatal("Wrong value on field 'BlockIndex'")
	}

	if block.Height != 0 {
		t.Fatal("Wrong value on field 'Height'")
	}

	if len(block.Tx) != 1 {
		t.Fatal("Wrong count items on field 'Tx'")
	}

	tx0 := block.Tx[0]

	if tx0.LockTime != 0 {
		t.Fatal("Wrong value on field 'Tx[0].LockTime'")
	}

	if tx0.Ver != 1 {
		t.Fatal("Wrong value on field 'Tx[0].Ver'")
	}

	if tx0.Size != 204 {
		t.Fatal("Wrong value on field 'Tx[0].Size'")
	}

	if tx0.Weight != 816 {
		t.Fatal("Wrong value on field 'Tx[0].Weight'")
	}

	if tx0.Time != 1231006505 {
		t.Fatal("Wrong value on field 'Tx[0].Time'")
	}

	if tx0.TxIndex != 14849 {
		t.Fatal("Wrong value on field 'Tx[0].TxIndex'")
	}

	if tx0.VinSz != 1 {
		t.Fatal("Wrong value on field 'Tx[0].VinSz'")
	}

	if tx0.Hash != "4a5e1e4baab89f3a32518a88c31bc87f618f76673e2cc77ab2127b7afdeda33b" {
		t.Fatal("Wrong value on field 'Tx[0].Hash'")
	}

	if tx0.RelayedBy != "0.0.0.0" {
		t.Fatal("Wrong value on field 'Tx[0].RelayedBy'")
	}

	inputs0 := tx0.Inputs[0]

	if inputs0.Sequence != 4294967295 {
		t.Fatal("Wrong value on field 'Tx[0].Inputs[0].Sequence'")
	}

	inputScript := "04ffff001d0104455468652054696d65732030332f4a616e2f32303039204368616e63656c6c6"
	inputScript += "f72206f6e206272696e6b206f66207365636f6e64206261696c6f757420666f722062616e6b73"
	if inputs0.Script != inputScript {
		t.Fatal("Wrong value on field 'Tx[0].Inputs[0].Script'")
	}

	out0 := tx0.Out[0]

	if out0.Spent != false {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Spent'")
	}

	if out0.TxIndex != 14849 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].TxIndex'")
	}

	if out0.Type != 0 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Type'")
	}

	if out0.Addr != "1A1zP1eP5QGefi2DMPTfTL5SLmv7DivfNa" {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Addr'")
	}

	if out0.Value != 5000000000 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].Value'")
	}

	if out0.N != 0 {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].N'")
	}

	outScript := "4104678afdb0fe5548271967f1a67130b7105cd6a828e03909a67962e0ea1f61deb"
	outScript += "649f6bc3f4cef38c4f35504e51ec112de5c384df7ba0b8d578a4c702b6bf11d5fac"
	if out0.Script != outScript {
		t.Fatal("Wrong value on field 'Tx[0].Out[0].N'")
	}
}
