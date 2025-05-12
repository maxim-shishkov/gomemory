package main

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/maxim-shishkov/gomemory/internal/compute/parser"
	"github.com/maxim-shishkov/gomemory/internal/storage/engine"
	"go.uber.org/zap"
)

func main() {
	logger := zap.NewExample()

	compute := parser.NewCompute(logger)
	storage := engine.NewStorage()
	db := engine.NewDataBase(storage, logger)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("command: ")
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		query, err := compute.Parse(str)
		if err != nil {
			fmt.Println(err)
			continue
		}

		val, err := db.Query(context.Background(), query)
		if err != nil {
			fmt.Println(err)
		}

		if val != "" {
			fmt.Println(val)
		}

	}

}
