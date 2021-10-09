package main

// Imported the required packages

import (
	"fmt"
	"net/http"

	"./controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	fmt.Println("Server Started!")
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/users", uc.CreateUser)

	http.ListenAndServe("localhost:3307", r) // Used Port: 3307 for hosting
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil { //  Checking whether the connection is established or disconnected
		fmt.Println("Error Encountered!")
		panic(err)
	}

	return s
}
