package parser

import (
	"github.com/guneyin/ethparser/client"
	"github.com/guneyin/ethparser/storage"
	"log/slog"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

const (
	latestBlockNum = "latest"
)

var blockNum = latestBlockNum

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []client.Transaction
}

type TXParser struct {
	client  *client.EthClient
	storage storage.Storage

	idCounter atomic.Uint32
}

func New() Parser {
	return &TXParser{
		client:  client.NewEthClient(),
		storage: storage.NewMemoryStorage(),
	}
}

type TransactionList struct {
	mu    sync.Mutex
	items map[string]client.Transaction
}

func NewTransactionList() *TransactionList {
	return &TransactionList{
		items: make(map[string]client.Transaction),
	}
}

// TXParser

func (p *TXParser) GetCurrentBlock() int {
	num, err := p.getLatestBlockNumber()
	if err != nil {
		slog.Error("Error on GetCurrentBlock", "Error", err.Error())
		return 0
	}

	return hexToInt(num)
}

func (p *TXParser) Subscribe(address string) bool {
	key := storage.NewKey(storage.Subscribe, address)

	if val := p.storage.Get(key); val != nil {
		slog.Warn("Already subscribed", "address", address)
	} else {
		p.storage.Set(key, true)
	}
	return true
}

type txJob func()

func txWorker(jobs <-chan txJob, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		job()
	}
}

func (p *TXParser) GetTransactions(address string) []client.Transaction {
	key := storage.NewKey(storage.Subscribe, address)

	if val := p.storage.Get(key); val == nil {
		slog.Error("Subscription not found", "address", address)
		return nil
	}

	block, err := p.getBlock(blockNum)
	if err != nil {
		slog.Error("Block could not fetch", "Error", err.Error())
		return nil
	}

	buffer := 5
	jobCount := len(block.Transactions)
	workerCount := (jobCount / buffer) + 1

	jobs := make(chan txJob, jobCount)
	wg := sync.WaitGroup{}

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go txWorker(jobs, &wg)
	}

	tl := NewTransactionList()
	for _, txHash := range block.Transactions {
		job := func() {
			tx, err := p.getTransaction(txHash)
			if err != nil {
				slog.Error("Transaction could not fetch", "Error", err.Error())
				return
			}

			if (tx.From == address) || (tx.To == address) {
				tl.put(tx)
			}
		}

		jobs <- job
	}

	close(jobs)
	wg.Wait()

	return tl.Items()
}

func (p *TXParser) getLatestBlockNumber() (string, error) {
	var response string
	err := p.client.Call(&response, client.MethodBlockNumber)
	if err != nil {
		return "", err
	}
	return response, nil
}

func (p *TXParser) getBlock(num string) (*client.Block, error) {
	block := new(client.Block)
	err := p.client.Call(&block, client.MethodBlockByNumber, num, false)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (p *TXParser) getTransaction(hash string) (*client.Transaction, error) {
	ts := new(client.Transaction)
	err := p.client.Call(&ts, client.MethodTransactionByHash, hash)
	if err != nil {
		return nil, err
	}
	return ts, nil
}

// TransactionList

func (tl *TransactionList) put(tx *client.Transaction) {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	tl.items[tx.Hash] = *tx
}

func (tl *TransactionList) Items() []client.Transaction {
	list := make([]client.Transaction, len(tl.items))
	var i int
	for _, tx := range tl.items {
		list[i] = tx
		i++
	}
	return list
}

// utils

func hexToInt(hexStr string) int {
	hexStr = strings.Replace(hexStr, "0x", "", 1)
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		slog.Error("Failed to parse block number", "err", err)
		return 0
	}
	return int(num)
}
