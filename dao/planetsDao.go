package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"go-rest-api/models"
)

type PlanetDao struct {
	Server   string
	Database string
}

var db *mgo.Database


const (
	COLLECTION = "planet"
)

func (m *PlanetDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *PlanetDao) GetAll() ([]models.Planet, error) {
	var planets []models.Planet
	err := db.C(COLLECTION).Find(bson.M{}).All(&planets)
	return planets, err
}

func (m *PlanetDao) GetByID(id string) (models.Planet, error) {
	var planet models.Planet
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&planet)
	return planet, err
}

func (m *PlanetDao) Create(planet models.Planet) error {
	err := db.C(COLLECTION).Insert(&planet)
	return err
}

func (m *PlanetDao) Delete(id string) error {
	err := db.C(COLLECTION).RemoveId(bson.ObjectIdHex(id))
	return err
}

func (m *PlanetDao) Update(id string, planet models.Planet) error {
	err := db.C(COLLECTION).UpdateId(bson.ObjectIdHex(id), &planet)
	return err
}