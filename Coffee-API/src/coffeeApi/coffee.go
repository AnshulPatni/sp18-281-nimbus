

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
var mongodb_database = "cmpe_Coffees"
var mongodb_collection = "starbucks_Coffees"



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
	mx.HandleFunc("/Coffees", homepageHandlerCoffees(formatter)).Methods("GET")
	mx.HandleFunc("/inventoryCoffees", inventoryHandlerCoffees(formatter)).Methods("GET")
	mx.HandleFunc("/cartItemsCoffees", cartHandlerCoffees(formatter)).Methods("GET")
	mx.HandleFunc("/searchInventoryCoffees", searchInventoryHandlerCoffees(formatter)).Methods("GET")
	//Below - PUT : Status 0 - 1
	mx.HandleFunc("/addToCartCoffees", starbucksAddToCartHandlerCoffees(formatter)).Methods("PUT")
	//Below - PUT : Status 1 -0
	mx.HandleFunc("/processOrdersCoffees", starbucksProcessOrdersHandlerCoffees(formatter)).Methods("PUT")
	//Below - Increase Likes
	mx.HandleFunc("/likeCoffees", likeHandlerCoffees(formatter)).Methods("PUT")
	mx.HandleFunc("/neworder", lNewOrder(formatter)).Methods("POST")

}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}



// Ping Application
func homepageHandlerCoffees(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks Dessert!"})
	}
}


//API returns  inventory for Status 0 items  to populate menu -
func inventoryHandlerCoffees(formatter *render.Render) http.HandlerFunc {
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
func cartHandlerCoffees(formatter *render.Render) http.HandlerFunc {

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
func searchInventoryHandlerCoffees(formatter *render.Render) http.HandlerFunc {
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
func starbucksAddToCartHandlerCoffees(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//parsing request bpdy from client and storing in array of bsons
			var CoffeeArrays []bson.M
			json.NewDecoder(req.Body).Decode(&CoffeeArrays)
	    	//fmt.Println("Added items to cart: ", m.starbucks)

		//establishing session with DB
			session, err := mgo.Dial(mongodb_server)
	        if err != nil {
	                panic(err)
	        }
	        defer session.Close()
	        session.SetMode(mgo.Monotonic, true)
	        c := session.DB(mongodb_database).C(mongodb_collection)


	  		//this is just for testing purposes. It is a CoffeeArrays of the request body from the client.
	        // var CoffeeArrays []bson.M
	        // err = c.Find(bson.M{"status":0}).All(&CoffeeArrays)
	        // if err != nil {
	        //         log.Fatal(err)
	        // }

	        //traversing throuhg request body and updating each bson based on condition.
	        for _, element := range CoffeeArrays {
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
}

//API processes items in cart - Updates status from 1 to 0
func starbucksProcessOrdersHandlerCoffees(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		//parsing request bpdy from client and storing in array of bsons
			// var sim []bson.M
	  //   	json.NewDecoder(req.Body).Decode(&sim)
	    	// fmt.Println("Update Gumball Inventory To: ", m.CountGumballs)

		//establishing session with DB
			session, err := mgo.Dial(mongodb_server)
	        if err != nil {
	                panic(err)
	        }
	        defer session.Close()
	        session.SetMode(mgo.Monotonic, true)
	        c := session.DB(mongodb_database).C(mongodb_collection)


	  		//this is just for testing purposes. It is a CoffeeArrays of the request body from the client.
	        var sim []bson.M
	        err = c.Find(bson.M{"status":1}).All(&sim)
	        if err != nil {
	                log.Fatal(err)
	        }

	        //travering throuhg request body and updating each bson based on condition.
	        for _, ele := range sim {
	        	fmt.Println("Item status changing from 1 to 0 :", ele )
	        	query := bson.M{"status":1}
	        	change := bson.M{"$set": bson.M{ "status" : 0}}

	        	err = c.Update(query, change)
	        	if err != nil {
	                log.Fatal(err)
	        	}

	        }

	        var results []bson.M
	        err = c.Find(bson.M{"status":0}).All(&results)
	        if err != nil {
	                log.Fatal(err)
	        }

	        fmt.Println("Items processed:", results )
			formatter.JSON(w, http.StatusOK, results)

	}
}

func likeHandlerCoffees(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        fmt.Println("Adding the like by 1  :")
        query := bson.M{"likes"}
	    change := bson.M{"$set": bson.M{ "likes"  }}

	    err = c.Update(query, change)
	    if err != nil {
	          log.Fatal(err)
	    }

	}

}

func lNewOrder(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uuid,_ := uuid.NewV4()
    	var ord = order {
					Id: uuid.String(),
					OrderStatus: "Order Placed",
		}
		if orders == nil {
			orders = make(map[string]order)
		}
		orders[uuid.String()] = ord
		queue_send(uuid.String())
		fmt.Println( "Orders: ", orders )
		formatter.JSON(w, http.StatusOK, ord)
	}
}
