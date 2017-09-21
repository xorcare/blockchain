// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

type Chart struct {
	Status      string   `json:"status"`
	Name        string   `json:"name"`
	Unit        string   `json:"unit"`
	Period      string   `json:"period"`
	Description string   `json:"description"`
	Values      []*Value `json:"values"`
}

type ChartPools map[string]uint64

type Value struct {
	X uint64  `json:"x"`
	Y float64 `json:"y"`
}

func (c *Client) GetPools() (response *ChartPools, e error) {
	response = &ChartPools{}
	var path = "/pools?format=json"
	e = c.DoRequest(path, response)

	return
}

func (c *Client) GetChart(chartType string) (response *Chart, e error) {
	response = &Chart{}
	var path = "/charts/" + chartType + "?format=json"
	e = c.DoRequest(path, response)

	return
}

// CURRENCY STATISTICS
//
// https://blockchain.info/charts/total-bitcoins
// Bitcoins in circulation
// The total number of bitcoins that have already been mined; in other words,
// the current supply of bitcoins on the network.
//
// https://blockchain.info/charts/market-price
// Market Price (USD)
// Average USD market price across major bitcoin exchanges.
//
// https://blockchain.info/charts/market-cap
// Market Capitalization
// The total USD value of bitcoin supply in circulation, as calculated by the
// daily average market price across major exchanges.
//
// https://blockchain.info/charts/trade-volume
// USD Exchange Trade Volume
// The total USD value of trading volume on major bitcoin exchanges.

// BLOCK DETAILS
//
// https://blockchain.info/charts/blocks-size
// Blockchain Size
// The total size of all block headers and transactions. Not including
// database indexes.
//
// https://blockchain.info/charts/avg-block-size
// Average Block Size
// The average block size in MB.
//
// https://blockchain.info/charts/n-orphaned-blocks
// Number Of Orphaned Blocks
// The total number of blocks mined but ultimately not attached to the main
// Bitcoin blockchain.
//
// https://blockchain.info/charts/n-transactions-per-block
// Average Number Of Transactions Per Block
// The average number of transactions per block.
//
// https://blockchain.info/charts/median-confirmation-time
// Median Confirmation Time
// The median time for a transaction to be accepted into a mined block and
// added to the public ledger.
//
// https://blockchain.info/charts/bip-9-segwit
// Percentage of blocks signalling SegWit support
//
// https://blockchain.info/charts/bitcoin-unlimited-share
// Percentage of blocks signalling Bitcoin Unlimited support
//
// https://blockchain.info/charts/nya-support
// New York Agreement support
// Percentage of blocks signalling for the New York Agreement over the last
// 200 blocks

// MINING INFORMATION
//
// https://blockchain.info/charts/hash-rate
// Hash Rate
// The estimated number of tera hashes per second the Bitcoin network is
// performing.
//
// https://blockchain.info/pools
// Hashrate Distribution
// An estimation of hashrate distribution amongst the largest mining pools
//
// https://blockchain.info/charts/difficulty
// Difficulty
// A relative measure of how difficult it is to find a new block.
//
// https://blockchain.info/charts/miners-revenue
// Mining Revenue
// Total value of coinbase block rewards and transaction fees paid to miners.
//
// https://blockchain.info/charts/transaction-fees
// Total Transaction Fees
// The total value of all transaction fees paid to miners (not including the
// coinbase value of block rewards).
//
// https://blockchain.info/charts/cost-per-transaction-percent
// Cost % of Transaction Volume
// A chart showing miners revenue as percentage of the transaction volume.
//
// https://blockchain.info/charts/cost-per-transaction
// Cost per Transaction
// A chart showing miners revenue divided by the number of transactions.

// BLOCKCHAIN WALLET ACTIVITY
//
// https://blockchain.info/charts/my-wallet-n-users
// Blockchain Wallet Users
