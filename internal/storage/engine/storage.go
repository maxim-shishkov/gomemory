package engine

import (
	"context"
	"fmt"
	"sync"
)

type Storage struct {
	memory map[string]string
	mx     *sync.Mutex
}

func NewStorage() *Storage {
	return &Storage{
		memory: make(map[string]string),
		mx:     &sync.Mutex{},
	}
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.memory[key] = value

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	s.mx.Lock()
	defer s.mx.Unlock()

	value, ok := s.memory[key]
	if !ok {
		return "", fmt.Errorf("not found")
	}

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	delete(s.memory, key)

	return nil
}
