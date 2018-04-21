package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	//"github.com/streadway/amqp"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	//"github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "cmpe_coffee"
var mongodb_collection = "starbucks_Coffee"



// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}



// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/coffee", homepageHandlerCoffee(formatter)).Methods("GET")
	mx.HandleFunc("/inventoryCoffee", inventoryHandlerCoffee(formatter)).Methods("GET")
	mx.HandleFunc("/cartItemsCoffee", cartHandlerCoffee(formatter)).Methods("GET")

}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}



// Ping Application
func homepageHandlerCoffee(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks Coffee!"})
	}
}


//API returns  inventory for Status 0 items  to populate menu -
func inventoryHandlerCoffee(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {
		//establishing session with DB
		fmt.Println("Inventory details:")
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

		var result []*bson.M

        err = c.Find(bson.M{"items":"mocha"}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)

	}
}


//API returns  inventory for Status 1 items  to populate cart -
func cartHandlerCoffee(formatter *render.Render) http.HandlerFunc{

	return func(w http.ResponseWriter, req *http.Request) {

		//establishing session with DB
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

		var result []*bson.M

        err = c.Find(bson.M{"status":1}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)



	}
}
