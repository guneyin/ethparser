package main

import (
	"context"
	"github.com/guneyin/ethparser/testutils"
	"net/http"
	"net/http/httptest"

	"testing"
	"time"
)

func TestMiddleware_startServer(t *testing.T) {
	ctx, done := context.WithTimeout(context.Background(), time.Second)
	defer done()

	errChan := make(chan error)

	go func() {
		errChan <- startServer(3000)
	}()

	select {
	case err := <-errChan:
		testutils.ShouldBeNil(t, err)
	case <-ctx.Done():
		return
	}
}

func TestMiddleware_ServeHTTP(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	mw := NewMiddleware(mux)
	mw.ServeHTTP(httptest.NewRecorder(), httptest.NewRequest(http.MethodGet, "/test", nil))

	httpServer := httptest.NewServer(mw)
	defer httpServer.Close()

	resp, err := http.Get(httpServer.URL + "/test")
	testutils.ShouldBeNil(t, err)
	testutils.ShouldBeEqual(t, http.StatusOK, resp.StatusCode)
}
