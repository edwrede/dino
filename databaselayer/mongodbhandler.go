package databaselayer

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MongodbHandler struct {
	*mgo.Session
}

func NewMongodbHandler(connection string) (*MongodbHandler, error) {
	session, err := mgo.Dial(connection)
	return &MongodbHandler{
		Session: session,
	}, err
}

func (handler *MongodbHandler) GetAvailableDinos() ([]Animal, error) {
	session := handler.getFreshSession()
	defer session.Close()
	animals := []Animal{}
	err := session.DB("dino").C("animals").Find(nil).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) GetDinoByNickname(nickname string) (Animal, error) {
	session := handler.getFreshSession()
	defer session.Close()
	a := Animal{}
	err := session.DB("dino").C("animals").Find(bson.M{"animal_nickname": nickname}).One(&a)
	return a, err
}

func (handler *MongodbHandler) GetDinosByType(dinoType string) ([]Animal, error) {
	session := handler.getFreshSession()
	defer session.Close()
	animals := []Animal{}
	err := session.DB("dino").C("animals").Find(bson.M{"animal_type": dinoType}).All(&animals)
	return animals, err
}

func (handler *MongodbHandler) AddAnimal(a Animal) error {
	session := handler.getFreshSession()
	defer session.Close()
	return session.DB("dino").C("animals").Insert(a)
}

func (handler *MongodbHandler) UpdateAnimal(a Animal, nickname string) error {
	session := handler.getFreshSession()
	defer session.Close()
	return session.DB("dino").C("animals").Update(bson.M{"animal_nickname": nickname}, a)
}

// When one does multi-threaded work in mongo you want to have copies of sessions making things multi-threaded safe
func (handler *MongodbHandler) getFreshSession() *mgo.Session {
	return handler.Session.Copy()
}
