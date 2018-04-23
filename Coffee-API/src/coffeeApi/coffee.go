

/*
	Starbucks API in Go

*/

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
var mongodb_database = "cmpe_Coffee"
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
	mx.HandleFunc("/Coffee", homepageHandlerCoffee(formatter)).Methods("GET")
	mx.HandleFunc("/inventoryCoffee", inventoryHandlerCoffee(formatter)).Methods("GET")
	mx.HandleFunc("/cartItemsCoffee", cartHandlerCoffee(formatter)).Methods("GET")
	mx.HandleFunc("/searchInventoryCoffee", searchInventoryHandlerCoffee(formatter)).Methods("GET")
	//Below - PUT : Status 0 - 1
	mx.HandleFunc("/addToCartCoffee", starbucksAddToCartHandlerCoffee(formatter)).Methods("PUT")
}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}



// Ping Application
func homepageHandlerCoffee(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks Dessert!"})
	}
}


//API returns  inventory for Status 0 items  to populate menu -
func inventoryHandlerCoffee(formatter *render.Render) http.HandlerFunc {
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

        err = c.Find(bson.M{"status":0}).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)

	}
}


//API returns  inventory for Status 1 items  to populate cart -
func cartHandlerCoffee(formatter *render.Render) http.HandlerFunc {

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


// API seacrhes inventory for a specific item and returns it
func searchInventoryHandlerCoffee(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M

        err = c.Find(bson.M{"item" : "mocha"}).One(&result)

        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Inventory details:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}



//API adds item to cart - Updates status from 0 to 1
func starbucksAddToCartHandlerCoffee(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//parsing request bpdy from client and storing in array of bsons
			var simulation []bson.M
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


	  		//this is just for testing purposes. It is a simulation of the request body from the client.
	        // var simulation []bson.M
	        // err = c.Find(bson.M{"status":0}).All(&simulation)
	        // if err != nil {
	        //         log.Fatal(err)
	        // }

	        //traversing throuhg request body and updating each bson based on condition.
	        for _, element := range simulation {
	        	fmt.Println("Item status changing from 0 to 1  :", element )
	        	query := bson.M{"status":0}
	        	change := bson.M{"$set": bson.M{ "status" : 1}}

	        	err = c.Update(query, change)
	        	if err != nil {
	                log.Fatal(err)
	        	}

	   		 }

	   		//this is our result
	        var results []bson.M
	        err = c.Find(bson.M{"status":1}).All(&results)
	        if err != nil {
	                log.Fatal(err)
	        }

	        fmt.Println("Items added to cart:", results )
			formatter.JSON(w, http.StatusOK, results)


	}
