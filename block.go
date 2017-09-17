package blockchain

type Block struct {
	Hash       string `json:"hash"`
	Ver        uint64 `json:"ver"`
	PrevBlock  string `json:"prev_block"`
	MrklRoot   string `json:"mrkl_root"`
	Time       uint64 `json:"time"`
	Bits       uint64 `json:"bits"`
	Fee        uint64 `json:"fee"`
	Nonce      uint64 `json:"nonce"`
	NTx        uint64 `json:"n_tx"`
	Size       uint64 `json:"size"`
	BlockIndex uint64 `json:"block_index"`
	MainChain  bool   `json:"main_chain"`
	Height     uint64 `json:"height"`
	Tx         []*Tx  `json:"tx"`
}

type LatestBlock struct {
	Hash       string   `json:"hash"`
	Time       uint64   `json:"time"`
	BlockIndex uint64   `json:"block_index"`
	Height     uint64   `json:"height"`
	TxIndexes  []uint64 `json:"txIndexes"`
}

type BlockHeight struct {
	Blocks []*Block `json:"blocks"`
}

func (c *Client) GetBlock(block string) (response *Block, e error) {
	response = &Block{}
	var path string = "/rawblock/" + block
	e = c.doRequest(path, response)

	return
}

func (c *Client) GetBlockHeight(blockHeight string) (response *BlockHeight, e error) {
	response = &BlockHeight{}
	var path string = "/block-height/" + blockHeight + "?format=json"
	e = c.doRequest(path, response)

	return
}

func (c *Client) GetLatestBlock() (response *LatestBlock, e error) {
	response = &LatestBlock{}
	var path string = "/latestblock"
	e = c.doRequest(path, response)

	return
}
