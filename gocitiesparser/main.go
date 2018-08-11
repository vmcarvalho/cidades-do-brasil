package main

import (
	"errors"
	"fmt"
	"github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/parser"
	"os"
)

func main() {
	args := os.Args[1:]
	err := parse(args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func parse(args []string) error {
	if len(args) < 1 {
		return errors.New("No input file argument")
	}
	inputFile := args[0]
	parser := parser.NewCitiesParser(inputFile)
	err := parser.Parse()
	if err != nil {
		return err
	}
	jsonStr, err := parser.ToJson()
	if err != nil {
		return err
	}
	fmt.Println(jsonStr)
	return nil
}
