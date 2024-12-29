package client

import "encoding/json"

type RPCRequest struct {
	Jsonrpc string    `json:"jsonrpc"`
	Method  EthMethod `json:"method"`
	Params  []any     `json:"params"`
	Id      uint32    `json:"id"`
}

type RPCResponse struct {
	Id      int             `json:"id"`
	Jsonrpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    string `json:"data"`
	} `json:"error"`
}

type Block struct {
	BaseFeePerGas         string        `json:"baseFeePerGas"`
	BlobGasUsed           string        `json:"blobGasUsed"`
	Difficulty            string        `json:"difficulty"`
	ExcessBlobGas         string        `json:"excessBlobGas"`
	ExtraData             string        `json:"extraData"`
	GasLimit              string        `json:"gasLimit"`
	GasUsed               string        `json:"gasUsed"`
	Hash                  string        `json:"hash"`
	LogsBloom             string        `json:"logsBloom"`
	Miner                 string        `json:"miner"`
	MixHash               string        `json:"mixHash"`
	Nonce                 string        `json:"nonce"`
	Number                string        `json:"number"`
	ParentBeaconBlockRoot string        `json:"parentBeaconBlockRoot"`
	ParentHash            string        `json:"parentHash"`
	ReceiptsRoot          string        `json:"receiptsRoot"`
	Sha3Uncles            string        `json:"sha3Uncles"`
	Size                  string        `json:"size"`
	StateRoot             string        `json:"stateRoot"`
	Timestamp             string        `json:"timestamp"`
	TotalDifficulty       string        `json:"totalDifficulty"`
	Transactions          []string      `json:"transactions"`
	TransactionsRoot      string        `json:"transactionsRoot"`
	Uncles                []interface{} `json:"uncles"`
	Withdrawals           []struct {
		Address        string `json:"address"`
		Amount         string `json:"amount"`
		Index          string `json:"index"`
		ValidatorIndex string `json:"validatorIndex"`
	} `json:"withdrawals"`
	WithdrawalsRoot string `json:"withdrawalsRoot"`
}

type Transaction struct {
	BlockHash        string `json:"blockHash"`
	BlockNumber      string `json:"blockNumber"`
	From             string `json:"from"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	Nonce            string `json:"nonce"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	Value            string `json:"value"`
	Type             string `json:"type"`
	ChainId          string `json:"chainId"`
	V                string `json:"v"`
	R                string `json:"r"`
	S                string `json:"s"`
}
