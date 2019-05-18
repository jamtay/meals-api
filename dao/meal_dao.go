package dao

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"meals-api/models"
)

type MealDao struct {
	Server string
	Database string
}

var mealDb *mgo.Database

const (
	MealCollection = "meals"
)

func (m *MealDao) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	mealDb = session.DB(m.Database)
}

func (m *MealDao) FindAll() ([]models.Meal, error){
	var meals []models.Meal
	err := mealDb.C(MealCollection).Find(bson.M{}).All(&meals)
	return meals, err
}

func (m *MealDao) Insert(meal models.Meal) error {
	err := mealDb.C(MealCollection).Insert(&meal)
	return err
}
