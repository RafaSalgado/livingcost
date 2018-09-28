package models

import "gopkg.in/mgo.v2/bson"

// Represents a livingcost, we uses bson keyword to tell the mgo driver how to name
// the properties in mongodb document
type Livingcost struct {
	ID              bson.ObjectId `bson:"_id" json:"id"`
	Zone            string        `bson:"zone" json:"zone"`
	Stratification  int           `bson:"stratification" json:"stratification"`
	Locality        string        `bson:"locality" json:"locality"`
	Costm2          int           `bson:"costm2" json:"costm2"`
	Costbasketgoods int           `bson:"costbasketgoods" json:"costbasketgoods"`
}
