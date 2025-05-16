package engine

import (
	"context"
	"fmt"
)

type Storage struct {
	memory map[string]string
}

func NewStorage() *Storage {
	return &Storage{
		memory: make(map[string]string),
	}
}

func (s *Storage) Set(ctx context.Context, key, value string) error {
	s.memory[key] = value

	return nil
}

func (s *Storage) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.memory[key]
	if !ok {
		return "", fmt.Errorf("not found")
	}

	return value, nil
}

func (s *Storage) Del(ctx context.Context, key string) error {
	delete(s.memory, key)

	return nil
}
