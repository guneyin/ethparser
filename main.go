package main

import (
	"flag"
	"fmt"
	"github.com/guneyin/ethparser/handler"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type Middleware struct {
	handler http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{handler}
}

func (m *Middleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	m.handler.ServeHTTP(w, r)
	slog.Info("HTTP Request", "method", r.Method, "path", r.URL.Path, "duration", time.Since(s).Round(time.Millisecond).String())
}

func main() {
	var port int
	flag.IntVar(&port, "port", 8080, "http port to listen on")
	flag.Parse()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	hnd := handler.New()

	mux := http.NewServeMux()
	mux.HandleFunc("/current-block", hnd.CurrentBlockHandler)
	mux.HandleFunc("/subscribe", hnd.SubscribeHandler)
	mux.HandleFunc("/transactions", hnd.TransactionsHandler)
	mw := NewMiddleware(mux)

	log.Printf("server running on http://127.0.0.1:%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mw))
}
