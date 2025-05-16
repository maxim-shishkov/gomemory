package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/maxim-shishkov/gomemory/internal/compute/parser"
	"go.uber.org/zap"
)

type Storer interface {
	Set(ctx context.Context, key, value string) error
	Get(ctx context.Context, key string) (string, error)
	Del(ctx context.Context, key string) error
}

func (s *Storage) Query() error {
	return nil
}

type DataBase struct {
	logger  *zap.Logger
	storage Storer
}

func NewDataBase(storage Storer, logger *zap.Logger) *DataBase {
	return &DataBase{
		storage: storage,
		logger:  logger,
	}
}

func (d *DataBase) Query(ctx context.Context, query parser.Query) (string, error) {
	args := query.Args()

	if len(args) == 0 {
		return "", fmt.Errorf("no arguments provided")
	}

	var val string
	var err error
	key := args[0]

	switch query.Command() {
	case parser.CommandSET:
		value := strings.Join(query.Args()[1:], " ")
		err = d.storage.Set(ctx, key, value)
	case parser.CommandGET:
		val, err = d.storage.Get(ctx, key)
	case parser.CommandDEL:
		err = d.storage.Del(ctx, key)
	default:
		d.logger.Error("unknown command")
		return "", fmt.Errorf("unknown command")
	}

	if err != nil {
		return "", fmt.Errorf("query: %w", err)
	}

	return val, nil
}
