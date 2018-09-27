package dao

import (
	"log"

	. "github.com/RafaSalgado/livingcost/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LivingcostsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database



const (
	COLLECTION = "livingcosts"
	hosts      = "dockercompose_mongodb_1:27017"
	database   = "db"
	username   = ""
	password   = ""
	collection = "jobs"
)
// info := &mgo.DialInfo{
// 		Addrs:    []string{hosts},
// 		Timeout:  60 * time.Second,
// 		Database: database,
// 		Username: username,
// 		Password: password,
// 	}

// Establish a connection to database
func (m *LivingcostsDAO) Connect() {
	info := &mgo.DialInfo{
		Addrs:    []string{hosts},
		Timeout:  60 * time.Second,
		Database: database,
		Username: username,
		Password: password,
	}
	session, err := mgo.DialWithInfo(info)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)

}

// Find list of livingcosts
func (m *LivingcostsDAO) FindAll() ([]Livingcost, error) {
	var livingcosts []Livingcost
	err := db.C(COLLECTION).Find(bson.M{}).All(&livingcosts)
	return livingcosts, err
}

// Find a Livingcost by its id
func (m *LivingcostsDAO) FindById(id string) (Livingcost, error) {
	var Livingcost Livingcost
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&Livingcost)
	return Livingcost, err
}

// Insert a Livingcost into database
func (m *LivingcostsDAO) Insert(Livingcost Livingcost) error {
	err := db.C(COLLECTION).Insert(&Livingcost)
	return err
}

// Delete an existing Livingcost
func (m *LivingcostsDAO) Delete(Livingcost Livingcost) error {
	err := db.C(COLLECTION).Remove(&Livingcost)
	return err
}

// Update an existing Livingcost
func (m *LivingcostsDAO) Update(Livingcost Livingcost) error {
	err := db.C(COLLECTION).UpdateId(Livingcost.ID, &Livingcost)
	return err
}
