/*
	CmpE281 - Starbucks - Backend APIs

*/

package main

import (
	"strconv"
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
var mongodb_database = "cmpe281"
var mongodb_collection = "starbucksSmoothies"




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

	mx.HandleFunc("/inventorySmoothies", inventoryHandlerSmoothies(formatter)).Methods("GET")
	mx.HandleFunc("/searchInventorySmoothies", searchInventoryHandlerSmoothies(formatter)).Methods("GET")
	mx.HandleFunc("/addToCartSmoothies", starbucksAddToCartHandlerSmoothies(formatter)).Methods("PUT")
	mx.HandleFunc("/cartItemsSmoothies", cartHandlerSmoothies(formatter)).Methods("GET")
		//Below - PUT : Status 1 -0 
	mx.HandleFunc("/processOrdersSmoothies", starbucksProcessOrdersHandlerSmoothies(formatter)).Methods("PUT")
	//Below - Increase Likes
	mx.HandleFunc("/likeIncreaseSmoothies/{item},{likes}", likeHandlerSmoothies(formatter)).Methods("PUT")

}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}




//API returns  inventory for Status 0 items  to populate menu - 
func inventoryHandlerSmoothies(formatter *render.Render) http.HandlerFunc {


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

        err = c.Find(nil).All(&result)

        fmt.Println("Inventory details:", result )
        formatter.JSON(w, http.StatusOK, result)


	}
}





// API seacrhes inventory for a specific item and returns it
func searchInventoryHandlerSmoothies(formatter *render.Render) http.HandlerFunc {
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
func starbucksAddToCartHandlerSmoothies(formatter *render.Render) http.HandlerFunc {
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

	        //traversing through request body and updating each bson based on condition.
	   		 for _, element := range simulation {
	   		 	for key, value := range element {

    				fmt.Println(key, value)
    				if(key == "item"){
    					fmt.Println("Retrieved item ", value)

    					query := bson.M{"item":value}
		        		change := bson.M{"$set": bson.M{ "status" : 1}}

			        	err = c.Update(query, change)
			        	if err != nil {
			                log.Fatal(err)
		        		}
    				}


  				}


	   		 }

	   		//this is our result
	        var finalresults []bson.M
	        err = c.Find(bson.M{"status":1}).All(&finalresults)
	        if err != nil {
	                log.Fatal(err)
	        } 
	       	   
	        fmt.Println("Items added to cart:", finalresults )
			formatter.JSON(w, http.StatusOK, finalresults)
		
	}
}

//API returns  inventory for Status 1 items  to populate cart - 
func cartHandlerSmoothies(formatter *render.Render) http.HandlerFunc {


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

//API processes items in cart - Updates status from 1 to 0
func starbucksProcessOrdersHandlerSmoothies(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {		


		//establishing session with DB
			session, err := mgo.Dial(mongodb_server)
	        if err != nil {
	                panic(err)
	        }
	        defer session.Close()
	        session.SetMode(mgo.Monotonic, true)
	        c := session.DB(mongodb_database).C(mongodb_collection)
	  		

	  		//this is just for testing purposes. It is a simulation of the request body from the client.
	        var sim []bson.M
	        err = c.Find(bson.M{ "status" : 1}).All(&sim)
	        if err != nil {
	                log.Fatal(err)
	        } 

	        //travering throuhg request body and updating each bson based on condition.
	        for _, ele := range sim {
	        	fmt.Println("Item status changing from 1 to 0 :", ele )
	        	query := bson.M{"status" : 1}
	        	change := bson.M{"$set": bson.M{ "status" : 0}}

	        	err = c.Update(query, change)
	        	if err != nil {
	                log.Fatal(err)
	        	}	

	        }	 
	       	   
	        fmt.Println("Items processed:", sim )
			formatter.JSON(w, http.StatusOK, sim)	
		
	}
}


//API takes in current number of likes and item liked -> returns updated like count(like+1)
func likeHandlerSmoothies(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {


		//parsing request string for value of current likes 
		params := mux.Vars(req)
		var likes string = params["likes"]
		var item string = params["item"]
		fmt.Println( "Item liked: ", item )
		fmt.Println( "Total likes: ", likes )


		

		session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)

        //converting likes - string to int
        i,err := strconv.Atoi(likes)
        if err != nil {
	          log.Fatal(err)
	    }	

        query := bson.M{"item" : item}
	    change := bson.M{"$set": bson.M{ "likes" : i+1 }}

	    err = c.Update(query, change)
	    if err != nil {
	          log.Fatal(err)
	    }	

	}	

}