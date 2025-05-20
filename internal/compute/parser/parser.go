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

	command := fields[0]
	args := fields[1:]

	switch command {
	case CommandSET, CommandGET, CommandDEL:
		return Query{
			command: command,
			args:    args,
		}, nil
	default:
		return Query{}, fmt.Errorf("unknown command")
	}
}
