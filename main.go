package main

import (
	"encoding/json"
	"log"
	"net/http"

	. "github.com/RafaSalgado/livingcost/config"
	. "github.com/RafaSalgado/livingcost/dao"
	. "github.com/RafaSalgado/livingcost/models"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = LivingcostsDAO{}

// GET list of livingcosts
func AllLivingcostsEndPoint(w http.ResponseWriter, r *http.Request) {
	livingcosts, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, livingcosts)
}

// GET a livingcost by its ID
func FindLivingcostEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}
	respondWithJson(w, http.StatusOK, livingcost)
}
func FindLivingcostByZoneEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindByZone(params["zone"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "400 Invalid Livingcost zone")
		return
	}
	respondWithJson(w, http.StatusOK, livingcost)
}


func FindLivingcostByLocalityEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindByLocality(params["locality"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost Locality")
		return
	}
	respondWithJson(w, http.StatusOK, livingcost)
}



// POST a new livingcost
func CreateLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var livingcost Livingcost
	if err := json.NewDecoder(r.Body).Decode(&livingcost); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	livingcost.ID = bson.NewObjectId()
	if err := dao.Insert(livingcost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, livingcost)
}

// PUT update an existing livingcost
func UpdateLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&livingcost); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(livingcost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

// DELETE an existing livingcost
func DeleteLivingcostEndPoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	livingcost, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Livingcost ID")
		return
	}

	if err := dao.Delete(livingcost); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

// Define HTTP request routes
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/livingcosts", AllLivingcostsEndPoint).Methods("GET")
	r.HandleFunc("/livingcosts", CreateLivingcostEndPoint).Methods("POST")
	r.HandleFunc("/livingcosts/{id}", UpdateLivingcostEndPoint).Methods("PUT")
	r.HandleFunc("/livingcosts/{id}", DeleteLivingcostEndPoint).Methods("DELETE")
	r.HandleFunc("/livingcosts/{id}", FindLivingcostEndpoint).Methods("GET")
	r.HandleFunc("/livingcosts/zone/{zone}", FindLivingcostByZoneEndpoint).Methods("GET")
	r.HandleFunc("/livingcosts/locality/{locality}", FindLivingcostByLocalityEndpoint).Methods("GET")
	if err := http.ListenAndServe(":4001", r); err != nil {
		log.Fatal(err)
	}
}
