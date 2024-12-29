package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

func GetAddrFromQuery(r *http.Request) (string, error) {
	addr := r.URL.Query().Get("addr")
	if strings.TrimSpace(addr) == "" {
		return "", errors.New("invalid address")
	}
	return addr, nil
}

func NewHttpResponseData(v any) []byte {
	data := make(map[string]any)
	data["result"] = v

	body, _ := json.Marshal(data)
	return body
}

func HttpOK(w http.ResponseWriter, body []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(body)
}

func HTTPError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
