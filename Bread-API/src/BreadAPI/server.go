/*
	Breads API in GO

*/

package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
)


// MongoDB Config
var mongodb_server = "localhost:27017"
var mongodb_database = "cmpe_breads"
var mongodb_collection = "starbucks_breads"



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
	mx.HandleFunc("/breads", homepageHandlerBreads(formatter)).Methods("GET")
}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}



// Ping Application
func homepageHandlerBreads(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks Breads!"})
	}
}
//returns inventory for Status 0
func inventoryHandlerBreads(formatter *render.Render) http.HandlerFunc{
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

        err = c.Find(bson.M{"items":"baguette"}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)

	}
}


//returns inventory for Status 1
func cartHandlerBreads(formatter *render.Render) http.HandlerFunc{

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
//Add to cart
func starbucksAddToCartHandlerBreads(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		
			var sim []bson.M
			json.NewDecoder(req.Body).Decode(&simulation)		
	    	//fmt.Println("Added items to cart: ", m.starbucks)

		//establishing session with DB
			session, err := mgo.Dial(mongodb_server)
	        if err != nil {
	                panic(err)
	        }
	        defer session.Close()
	        session.SetMode(mgo.Monotonic, true)
	        c := session.DB(mongodb_database).C(mongodb_collection)
	  		
	        var simulation []bson.M
	        err = c.Find(bson.M{"status":0}).All(&sim)
	        if err != nil {
			log.Fatal(err)
			
	        } 

		formatter.JSON(w, http.StatusOK, sim)
	}