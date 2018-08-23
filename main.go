package main

import (
	"errors"
	"fmt"
	database "github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/database"
	model "github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/model"
	"github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/parser"
	"os"
	"strings"
)

func main() {
	args := os.Args
	fmt.Println("Starting Go Parser")

	err := checkArgs(args)
	if err != nil {
		panic(err)
	}
	fmt.Println("... parsing input")
	cities, err := parseInputCsv(args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("... dumping to database ", args[2])
	db := databaseAdapterFactory(args[2])
	err = dumpToDatabase(db, cities)
	if err != nil {
		panic(err)
	}
	fmt.Println("... Finished!")
}

// factory method for the chosen database.DatabaseAdapter implementation
func databaseAdapterFactory(url string) database.DatabaseAdapter {
	return database.NewMongoDatabase(url)
}

func checkArgs(args []string) error {
	if len(args) < 2 {
		msg := fmt.Sprintf("Missing args.\nUsage:\t%s input_file database_url\n", args[0])
		msg = msg + fmt.Sprintf("Cmd:\t%s\n", strings.Join(args[:], " "))
		return errors.New(msg)
	}
	return nil
}

// Parses the csv input file in path 'inputFile' to an array of model.City struct
func parseInputCsv(inputFile string) ([]model.City, error) {
	parser := parser.NewCitiesParser(inputFile)
	err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return parser.GetCities()
}

// Dump an array of model.City to a database using database.DatabaseAdapter interface
func dumpToDatabase(db database.DatabaseAdapter, cities []model.City) error {
	for _, city := range cities {
		err := db.Add(city)
		if err != nil {
			return err
		}
	}
	return nil
}
