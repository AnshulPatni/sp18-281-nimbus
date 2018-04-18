package main

import (
	"fmt"
	"log"
	"net/http"
	//"encoding/json"
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
var mongodb_database = "cmpe281"
var mongodb_collection = "starbucks"

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
	mx.HandleFunc("/", homepageHandler(formatter)).Methods("GET")
	mx.HandleFunc("/inventory", inventoryHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/starbucks", starbucksUpdateHandler(formatter)).Methods("PUT")
	// mx.HandleFunc("/order", starbucksNewOrderHandler(formatter)).Methods("POST")
	// mx.HandleFunc("/order/{id}", starbucksOrderStatusHandler(formatter)).Methods("GET")
	// mx.HandleFunc("/order", starbucksOrderStatusHandler(formatter)).Methods("GET")
}

// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

//"Id" : "1.0"

// Ping Application
func homepageHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.JSON(w, http.StatusOK, struct{ Test string }{"Welcome to Starbucks!"})
	}
}

// API returns entire inventory
func inventoryHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)
        var result bson.M
        err = c.Find(bson.M{}).One(&result)
        if err != nil {
                log.Fatal(err)
        }
        fmt.Println("Inventory details:", result )
		formatter.JSON(w, http.StatusOK, result)
	}
}


// API Create New starbucks Order
func starbucksNewOrderHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		uuid,_ := uuid.NewV4()
    	var ord = order {
//					Id: uuid.String(),            		
					status: 0,
		}
		if orders == nil {
			orders = make(map[string]order)
		}
		orders[uuid.String()] = ord
		fmt.Println( "Orders: ", orders )
		formatter.JSON(w, http.StatusOK, ord)
	}
}

// API Get Order Status
func starbucksOrderStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		params := mux.Vars(req)
		var uuid string = params["id"]
		fmt.Println( "Order ID: ", uuid )
		if uuid == ""  {
			fmt.Println( "Orders:", orders )
			var orders_array [] order
			for key, value := range orders {
    			fmt.Println("Key:", key, "Value:", value)
    			orders_array = append(orders_array, value)
			}
			formatter.JSON(w, http.StatusOK, orders_array)
		} else {
			var ord = orders[uuid]
			fmt.Println( "Order: ", ord )
			formatter.JSON(w, http.StatusOK, ord)
		}
	}
}

// API Process Orders 
func starbucksProcessOrdersHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		// Open MongoDB Session
		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

       	// Get starbucks Inventory 
        var result bson.M
        err = c.Find(bson.M{"item" : "mocha"}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

 		var count int = result["count"]
        fmt.Println("Current Inventory:", count )

		for i := 0; i < len(order_ids); i++ {
			var order_id = order_ids[i]
			fmt.Println("Order ID:", order_id)
			var ord = orders[order_id] 
			ord.OrderStatus = "Order Processed"
			orders[order_id] = ord
			count -= 1
		}
		fmt.Println( "Orders: ", orders , "New Inventory: ", count)

		// Update starbucks Inventory
		query := bson.M{"item" : "mocha"}
        change := bson.M{"$set": bson.M{ "status" : 0}}
        err = c.Update(query, change)
        if err != nil {
                log.Fatal(err)
        }

		// Return Order Status
		formatter.JSON(w, http.StatusOK, orders)
	}
}