package storage

import (
	"sync"
)

type MemoryStorage struct {
	mu      sync.Mutex
	storage map[string]any
}

func NewMemoryStorage() Storage {
	return &MemoryStorage{storage: make(map[string]any)}
}

func (s *MemoryStorage) Get(key string) any {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.storage[key]
}

func (s *MemoryStorage) Set(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.storage[key] = value
}
