/*
        Starbucks - Coffee API
*/

package main

import (
        "strconv"
        "fmt"
        "log"
        "net/http"
        "github.com/unrolled/render"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)


//var mongodb_server = "localhost:27017"
var mongodb_server = "mongodb://10.0.2.244:27017,10.0.2.106:27017,10.0.2.64:27017,10.0.5.10:27017,10.0.5.14:27017/?replicaSet=rs1"
var mongodb_database = "cmpe281"
var mongodb_collection = "starbucksCoffees"
var mongodb_collection2 = "store"

func homepageHandler(rw http.ResponseWriter, req *http.Request) {
        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Credentials", "true")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        }

        formatter := render.New(render.Options{
                            IndentJSON: true,
                        })

        formatter.JSON(rw, http.StatusOK, struct{ Test string }{"Welcome to Starbucks!"})

}


func inventoryHandlerCoffees(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Credentials","true")
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        }



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



                        formatter := render.New(render.Options{
                            IndentJSON: true,
                        })

                        fmt.Println("Inventory details:", result )
                        formatter.JSON(rw, http.StatusOK, result)

}

func cartHandlerCoffees(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }


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

        formatter := render.New(render.Options{
                            IndentJSON: true,
                        })

        fmt.Println("Inventory details:", result )
        formatter.JSON(rw, http.StatusOK, result)

}

func starbucksAddToCartHandlerCoffees(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }



    //io.WriteString(res, "\nphone: "+req.FormValue("phone"))
    fmt.Println( "In add to cart!" )
        //establishing session with DB
        session, err := mgo.Dial(mongodb_server)
          if err != nil {
                  panic(err)
          }
          defer session.Close()
          session.SetMode(mgo.Monotonic, true)
          c := session.DB(mongodb_database).C(mongodb_collection)



          query := bson.M{"item":req.FormValue("item")}
          change := bson.M{"$set": bson.M{ "status" : 1}}

          err = c.Update(query, change)
          if err != nil {
            log.Fatal(err)
          }



        //this is our result
          var finalresults []bson.M
          err = c.Find(bson.M{"status":1}).All(&finalresults)
          if err != nil {
                  log.Fatal(err)
          }

          formatter := render.New(render.Options{
                            IndentJSON: true,
                        })

          fmt.Println("Items added to cart:", finalresults )
          formatter.JSON(rw, http.StatusOK, finalresults)

}




func starbucksProcessOrdersHandlerCoffees(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }


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

                    formatter := render.New(render.Options{
                            IndentJSON: true,
                        })



          fmt.Println("Items processed:", sim )
          formatter.JSON(rw, http.StatusOK, sim)

}
