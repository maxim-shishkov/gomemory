package parser

import (
	"fmt"
	"strings"
)

const (
	CommandGET = "GET"
	CommandSET = "SET"
	CommandDEL = "DEL"
)

func (c *Compute) Parse(text string) (Query, error) {
	fields := strings.Fields(text)
	if len(fields) < 0 {
		return Query{}, fmt.Errorf("empty command")
	}

	var query Query

	switch fields[0] {
	case CommandSET:
		query = Query{
			command: CommandSET,
			args:    fields[1:],
		}
	case CommandGET:
		query = Query{
			command: CommandGET,
			args:    fields[1:],
		}
	case CommandDEL:
		query = Query{
			command: CommandDEL,
			args:    fields[1:],
		}
	default:
		return Query{}, fmt.Errorf("unknown command")
	}

	return query, nil
}
