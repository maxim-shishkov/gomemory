package engine

import (
	"context"
	"fmt"
	"strings"

	"github.com/maxim-shishkov/gomemory/internal/compute/parser"
	"go.uber.org/zap"
)

var _ Storer = &Storage{}

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
	var val string
	var err error

	switch query.Command() {
	case parser.CommandSET:
		arg := strings.Join(query.Args()[1:], " ")
		err = d.storage.Set(ctx, query.Args()[0], arg)
	case parser.CommandGET:
		arg := query.Args()[0]
		val, err = d.storage.Get(ctx, arg)
	case parser.CommandDEL:
		arg := query.Args()[0]
		err = d.storage.Del(ctx, arg)
	default:
		d.logger.Error("unknown command")
		return "", fmt.Errorf("unknown command")
	}

	if err != nil {
		return "", fmt.Errorf("query: %w", err)
	}

	return val, nil
}
