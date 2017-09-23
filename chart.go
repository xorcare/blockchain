// Copyright 2017 Vasilyuk Vasiliy. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package blockchain

//Chart API data struct
type Chart struct {
	Status      string   `json:"status"`
	Name        string   `json:"name"`
	Unit        string   `json:"unit"`
	Period      string   `json:"period"`
	Description string   `json:"description"`
	Values      []*Value `json:"values"`
}

// Pools information map
type ChartPools map[string]uint64

//Charts API values
type Value struct {
	X uint64  `json:"x"`
	Y float64 `json:"y"`
}

type Stats struct {
	MarketPriceUsd                float64 `json:"market_price_usd"`
	HashRate                      float64 `json:"hash_rate"`
	TotalFeesBtc                  int64   `json:"total_fees_btc"`
	NBtcMined                     int64   `json:"n_btc_mined"`
	NTx                           uint64  `json:"n_tx"`
	NBlocksMined                  uint64  `json:"n_blocks_mined"`
	MinutesBetweenBlocks          float64 `json:"minutes_between_blocks"`
	Totalbc                       int64   `json:"totalbc"`
	NBlocksTotal                  uint64  `json:"n_blocks_total"`
	EstimatedTransactionVolumeUsd float64 `json:"estimated_transaction_volume_usd"`
	BlocksSize                    uint64  `json:"blocks_size"`
	MinersRevenueUsd              float64 `json:"miners_revenue_usd"`
	Nextretarget                  int64   `json:"nextretarget"`
	Difficulty                    int64   `json:"difficulty"`
	EstimatedBtcSent              int64   `json:"estimated_btc_sent"`
	MinersRevenueBtc              int64   `json:"miners_revenue_btc"`
	TotalBtcSent                  int64   `json:"total_btc_sent"`
	TradeVolumeBtc                float64 `json:"trade_volume_btc"`
	TradeVolumeUsd                float64 `json:"trade_volume_usd"`
	Timestamp                     uint64  `json:"timestamp"`
}

// Stats API Get the data behind Blockchain's stats
// This method can be used to get the data behind Blockchain.info's stats.
// URL: https://blockchain.info/stats
func (c *Client) GetStats() (response *Stats, e error) {
	response = &Stats{}
	e = c.DoRequest("/stats", response, map[string]string{"format": "json"})

	return
}

// Pools API Get the data behind Blockchain's pools information
// This method can be used to get the data behind Blockchain.info's pools information.
// URL: https://blockchain.info/pools
func (c *Client) GetPools() (response *ChartPools, e error) {
	response = &ChartPools{}
	e = c.DoRequest("/pools", response, map[string]string{"format": "json"})

	return
}

// Charts API Get the data behind Blockchain's charts
// This method can be used to get and manipulate data behind all Blockchain.info's charts.
// URL: https://blockchain.info/charts
//
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
//
//
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
//
//
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
//
//
// BLOCKCHAIN WALLET ACTIVITY
//
// https://blockchain.info/charts/my-wallet-n-users
// Blockchain Wallet Users
func (c *Client) GetChart(chartType string, params ...map[string]string) (response *Chart, e error) {
	options := map[string]string{"format": "json"}
	if len(params) > 0 {
		for k, v := range params[0] {
			options[k] = v
		}
	}
	response = &Chart{}
	e = c.DoRequest("/charts/"+chartType, response, options)

	return
}
