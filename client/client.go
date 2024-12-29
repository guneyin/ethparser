package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"sync/atomic"
	"time"
)

type EthMethod string

const (
	ethApiUrl = "https://ethereum-rpc.publicnode.com"

	MethodBlockNumber       = EthMethod("eth_blockNumber")
	MethodBlockByNumber     = EthMethod("eth_getBlockByNumber")
	MethodTransactionByHash = EthMethod("eth_getTransactionByHash")
)

type EthClient struct {
	client    *http.Client
	idCounter atomic.Uint32
}

func NewEthClient() *EthClient {
	return &EthClient{
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (e *EthClient) Call(result any, method EthMethod, params ...any) error {
	if result != nil && reflect.TypeOf(result).Kind() != reflect.Ptr {
		return fmt.Errorf("incorrect result type: %v", result)
	}

	data := RPCRequest{
		Jsonrpc: "2.0",
		Method:  method,
		Params:  params,
		Id:      e.idCounter.Add(1),
	}
	body, _ := json.Marshal(data)

	req := &http.Request{}
	req.URL, _ = url.Parse(ethApiUrl)
	req.Method = http.MethodPost
	req.Body = io.NopCloser(bytes.NewBuffer(body))

	res, err := e.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return parseRPCResponse(res, result)
}

func parseRPCResponse(r *http.Response, obj any) error {
	rpc := new(RPCResponse)

	body, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(body, &rpc)
	if err != nil {
		return err
	}

	if rpc.Error.Code != 0 {
		return errors.New(rpc.Error.Message)
	}

	return json.Unmarshal(rpc.Result, obj)
}
