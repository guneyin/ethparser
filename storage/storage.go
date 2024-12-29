package storage

import "fmt"

type Object string

const (
	Subscribe = Object("SUBSCRIBE")
)

type Storage interface {
	Get(key string) any
	Set(key string, value any)
}

func NewKey(so Object, addr string) string {
	return fmt.Sprintf("%s:%s", so, addr)
}
