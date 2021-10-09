package controllers

// Imported the required packages

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../models"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Initiated the Mongo Session
type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// Function for GetUser request with specified Id

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.MongoField{}

	if err := uc.session.DB("First_Database").C("First COllection").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
	}
	fmt.Println("\n", u)
	uj, err := json.Marshal(u) // Used Json.Marshal to encode the given details
	if err != nil {
		fmt.Println(err)
	}
	// Coverted the return to Json object type
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(202)
	fmt.Fprint(w, "\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	u := models.MongoField{}

	json.NewDecoder(r.Body).Decode(&u)

	u.Id = string(bson.NewObjectId())

	uc.session.DB("First_Database").C("First COllection").Insert(u)

	uj, err := json.Marshal(u)

	if err != nil { // Checking if it thows any error
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "%s\n", uj) // Returned the Create User in MongoDb Compass
}
