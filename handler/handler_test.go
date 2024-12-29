package handler

import (
	"github.com/guneyin/ethparser/testutils"
	"net/http"
	"net/http/httptest"
	"testing"
)

var hnd *Handler

func init() {
	hnd = New()
}

const testAddress = "0x0457896"

func TestHandler_CurrentBlockHandler(t *testing.T) {
	resp := makeCall(hnd.CurrentBlockHandler, "/current-block")
	testutils.ShouldBeEqual(t, 200, resp.StatusCode)
}

func TestHandler_SubscribeHandler(t *testing.T) {
	resp := makeCall(hnd.SubscribeHandler, "/subscribe?addr="+testAddress)
	testutils.ShouldBeEqual(t, 200, resp.StatusCode)

	resp = makeCall(hnd.SubscribeHandler, "/subscribe?addr=")
	testutils.ShouldNotBeEqual(t, 200, resp.StatusCode)
}

func TestHandler_TransactionsHandler(t *testing.T) {
	resp := makeCall(hnd.TransactionsHandler, "/transactions?addr="+testAddress)
	testutils.ShouldBeEqual(t, 200, resp.StatusCode)

	resp = makeCall(hnd.TransactionsHandler, "/transactions?addr=")
	testutils.ShouldNotBeEqual(t, 200, resp.StatusCode)
}

func makeCall(handler http.HandlerFunc, uri string) *http.Response {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, uri, nil)
	handler(w, r)
	return w.Result()
}
