package engine

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStorage(t *testing.T) {
	ctx := context.Background()

	s := &Storage{
		memory: make(map[string]string),
	}
	key := "key"

	err := s.Set(ctx, key, "value")
	assert.NoError(t, err)

	val, err := s.Get(ctx, key)
	assert.NoError(t, err)
	assert.Equal(t, val, "value")

	err = s.Del(ctx, key)
	assert.NoError(t, err)

	val, err = s.Get(ctx, key)
	assert.Error(t, fmt.Errorf("not found"), err)
	assert.Equal(t, val, "")
}
