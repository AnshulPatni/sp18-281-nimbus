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
var mongodb_collection = "starbucksDesserts"




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

	mx.HandleFunc("/inventoryDesserts", inventoryHandlerDesserts(formatter)).Methods("GET")
	mx.HandleFunc("/cartItemsDesserts", cartHandlerDesserts(formatter)).Methods("GET")
	mx.HandleFunc("/searchInventoryDesserts", searchInventoryHandlerDesserts(formatter)).Methods("GET")
	//Below - PUT : Status 0 - 1
	mx.HandleFunc("/addToCartDesserts", starbucksAddToCartHandlerDesserts(formatter)).Methods("PUT")
	//Below - PUT : Status 1 -0 
	mx.HandleFunc("/processOrdersDesserts", starbucksProcessOrdersHandlerDesserts(formatter)).Methods("PUT")
	//Below - Increase Likes
	mx.HandleFunc("/likeIncreaseDesserts/{item},{likes}", likeHandlerDesserts(formatter)).Methods("PUT")

}



// Helper Functions
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}





//API returns  inventory for Status 0 items  to populate menu - 
func inventoryHandlerDesserts(formatter *render.Render) http.HandlerFunc {


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



//API returns  inventory for Status 1 items  to populate cart - 
func cartHandlerDesserts(formatter *render.Render) http.HandlerFunc {


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
func searchInventoryHandlerDesserts(formatter *render.Render) http.HandlerFunc {
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
func starbucksAddToCartHandlerDesserts(formatter *render.Render) http.HandlerFunc {
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

//API processes items in cart - Updates status from 1 to 0
func starbucksProcessOrdersHandlerDesserts(formatter *render.Render) http.HandlerFunc {
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
func likeHandlerDesserts(formatter *render.Render) http.HandlerFunc {
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





/*

	
	-- StarbucksDesserts MongoDB Create Database

		Database Name: cmpe281
		Collection Name: starbucksDesserts

  	-- StarbucksDesserts MongoDB Collection (Create Document) --
		
		//menu
	    db.starbucksDesserts.insert(
		    { 
		      typeService : 'dessert',
		      item: 'chocolate scone',
		      cost: NumberInt(5),
		      likes: NumberInt(0),
		      status : NumberInt(0)
		    }
		) ;
		    db.starbucksDesserts.insert(
		    { 
		      typeService : 'dessert',
		      item: 'pecan tart',
		      cost: NumberInt(4),
		      likes: NumberInt(0),
		      status : NumberInt(0)
		    }
		) ;


		db.starbucksDesserts.insert(
		    { 
		      typeService : 'dessert',
		      item: 'coffee cake',
		      cost: NumberInt(3),
		      likes: NumberInt(0),
		      status : NumberInt(0)
		    }
		) ;

	
    -- StarbucksDesserts Menu MongoDB Collection - Find Menu Documents --

    db.starbucksDesserts.find( { item: 'coffee' } ) ;
    db.starbucksDesserts.find( { item: 'mocha' } ) ;

    
    -- StarbucksDesserts Delete Menu Collection

    db.starbucksDesserts.remove({})



 */
