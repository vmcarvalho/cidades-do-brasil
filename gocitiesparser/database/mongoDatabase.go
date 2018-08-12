package database

import (
	"errors"
	driver "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	model "github.com/vmcarvalho/cidades-do-brasil/gocitiesparser/model"
)

const (
	COLLECTION = "City"
	NAME       = "name"
	UF         = "uf"
	LAT        = "lat"
	LON        = "long"
)

type MongoDatabase struct {
	url            string
	session        *driver.Session
	collectionName string
	collection     *driver.Collection
}

func NewMongoDatabase(url string) DatabaseAdapter {
	m := newInstance()
	m.url = url
	return m
}

func newInstance() *MongoDatabase {
	m := new(MongoDatabase)
	m.collectionName = COLLECTION
	return m
}

func (m *MongoDatabase) connect() error {
	var err error
	m.session, err = driver.Dial(m.url)
	if err != nil {
		return err
	}
	db := m.session.DB("")
	if db == nil {
		return errors.New("DB is nil!")
	}
	m.collection = m.session.DB("").C(m.collectionName)
	return err
}

func (m *MongoDatabase) disconnect() {
	if m.session == nil {
		return
	}
	m.session.Close()
}

func (m *MongoDatabase) setCollection(name string) {
	m.collectionName = name
}

func (m MongoDatabase) Add(city model.City) error {
	err := m.connect()
	if err != nil {
		return err
	}
	defer m.disconnect()

	err = m.collection.Insert(city)
	if err != nil {
		return err
	}

	return nil
}

func (m MongoDatabase) Remove(city model.City) (bool, error) {
	err := m.connect()
	if err != nil {
		return false, err
	}
	defer m.disconnect()

	var changeInfo *driver.ChangeInfo
	changeInfo, err = m.collection.RemoveAll(bson.M{NAME: city.Name, UF: city.Uf})
	if err != nil {
		return false, err
	}

	return (changeInfo.Removed > 0), nil
}

func (m MongoDatabase) List() ([]model.City, error) {
	err := m.connect()
	if err != nil {
		return nil, err
	}
	defer m.disconnect()

	var results []model.City
	err = m.collection.Find(nil).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (m MongoDatabase) SearchByName(cityName string) ([]model.City, error) {
	err := m.connect()
	if err != nil {
		return nil, err
	}
	defer m.disconnect()

	var results []model.City
	err = m.collection.Find(bson.M{NAME: cityName}).All(&results)
	if err != nil {
		return nil, err
	}

	return results, nil
}
