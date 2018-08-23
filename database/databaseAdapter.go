package database

import (
	model "github.com/vmcarvalho/cidades-do-brasil/model"
)

type DatabaseAdapter interface {
	Add(city model.City) error
	Remove(city model.City) (bool, error)
	List() ([]model.City, error)
	SearchByName(cityName string) ([]model.City, error)
}
