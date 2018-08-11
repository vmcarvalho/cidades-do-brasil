package parser

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	model "github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/model"
	"io"
	"os"
	"strconv"
)

type CitiesParser struct {
	inputFile   string
	citiesArray []model.City
}

const CATEGORY_INDEX = 2
const NAME_INDEX = 1
const UF_INDEX = 5
const LAT_INDEX = 6
const LON_INDEX = 7
const CITY_CATEGORY = "LIM"

func NewCitiesParser(input string) *CitiesParser {
	p := new(CitiesParser)
	p.inputFile = input
	return p
}

func (this *CitiesParser) Parse() error {
	csvFile, err := os.Open(this.inputFile)
	if err != nil {
		return err
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))
	reader.Comma = ';'
	this.citiesArray = []model.City{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if line[CATEGORY_INDEX] != CITY_CATEGORY {
			continue
		}

		lat, err := strconv.ParseFloat(line[LAT_INDEX], 64)
		if err != nil {
			return err
		}
		lon, err := strconv.ParseFloat(line[LON_INDEX], 64)
		if err != nil {
			return err
		}

		this.citiesArray = append(this.citiesArray, model.City{
			Name:      line[NAME_INDEX],
			Uf:        line[UF_INDEX],
			Latitude:  lat,
			Longitude: lon,
		})
	}
	return nil
}

func (this CitiesParser) GetCities() ([]model.City, error) {
	if this.citiesArray == nil {
		return nil, errors.New("citiesArray is null!")
	}
	return this.citiesArray, nil
}

func (this CitiesParser) ToJson() (string, error) {
	cities, err := this.GetCities()
	if err != nil {
		return "", err
	}
	byteArray, err := json.Marshal(cities)
	if err != nil {
		return "", err
	}
	return string(byteArray), nil
}
