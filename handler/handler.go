package handler

import (
	"fmt"
	"github.com/guneyin/ethparser/parser"
	"net/http"
)

type Handler struct {
	p parser.Parser
}

func New() *Handler {
	return &Handler{
		p: parser.New(),
	}
}

func (h *Handler) CurrentBlockHandler(w http.ResponseWriter, r *http.Request) {
	cb := h.p.GetCurrentBlock()
	HttpOK(w, NewHttpResponseData(cb))
}

func (h *Handler) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	addr, err := GetAddrFromQuery(r)
	if err != nil {
		HTTPError(w, err)
		return
	}

	ok := h.p.Subscribe(addr)
	if !ok {
		HTTPError(w, fmt.Errorf("subscribe failed"))
		return
	}

	HttpOK(w, NewHttpResponseData(ok))
}

func (h *Handler) TransactionsHandler(w http.ResponseWriter, r *http.Request) {
	addr, err := GetAddrFromQuery(r)
	if err != nil {
		HTTPError(w, err)
		return
	}

	tx := h.p.GetTransactions(addr)

	HttpOK(w, NewHttpResponseData(tx))
}
