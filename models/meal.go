package models

import "gopkg.in/mgo.v2/bson"

type Meal struct {
	ID	bson.ObjectId `bson:"_id" json:"id"`
	Name string `bson:"name" json:"name"`
	Description string `bson:"description	" json:"description"`
	OverallCost float32 `bson:"overall_cost" json:"overall_cost"`
}
