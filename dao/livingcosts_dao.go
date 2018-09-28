package dao

import (
	"log"
	"time"
	"math/rand"

	. "github.com/RafaSalgado/livingcost/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LivingcostsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

var livingcost Livingcost

const (
	COLLECTION = "livingcosts"
	hosts      = "localhost:27017"
	database   = "livingcosts1_db"
	username   = ""
	password   = ""
)

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
	fulldata()

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

func fulldata() {
	local := [20]string{"Bosa", "Kennedy", "Usaquen", "Chapinero", "Santa Fe", "San Cristobal", "Usme" ,"Tunjuelito", "Fontibon", "Engativa", "Suba",
		"Barrios Unidos", "Teusaquillo", "Los Marires", "Antonio Nariño", "Sumapaz", "Ciudad Bolivar" , "Rafael Uribe Uribe", "La candelaria",
	 	"Puente Aranda"}

	zonas := [34]strin{"Venecia" , "Cedritos" , "Santa Barbara", "Lijaca", "La Macarena", "El libertador", "Carvajal", "Madelena", "Marly",
											"Modelia", "La Jimenez", "Galerias","La castellana", "Polo club", "San Antonio", "Carbonel" , "Casa linda", "Meissen",
										 	"Lucero alto", "La Belleza", "Hayuelos", "El Dorado", "Ricauete", "Santa Isabel", "Salitre", "Sierra Morena",
											"Ciudad Bolivar", "Tunal", "Fatima", "Marsella", "Banderas", "Patio Bonito", "Aures", "Lisboa"}
	collection := db.C(COLLECTION)
	for j := 0; j <= 50; j++ {
		livingcost.ID = bson.NewObjectId()
		livingcost.zone = zonas[rand.Intn(34)]
		livingcost.stratification = rand.Intn(7)
		livingcost.locality = local[rand.Intn(20)]
		livingcost.costbasketgoods = rand.Intn( 7000000)
		livingcost.costbasketgoods = rand.Intn(400000)
		err := collection.Insert(livingcost)
		log.Println(j)
		if err != nil {
			log.Fatal(err)
		}
	}
}
