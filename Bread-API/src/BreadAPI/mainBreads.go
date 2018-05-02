/*
        Starbucks - Breads API
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
var mongodb_server = "mongodb://10.0.0.123:27017,10.0.0.199:27017,10.0.0.169:27017,10.0.2.233:27017,10.0.2.4:27017/?replicaSet=repSet"
var mongodb_database = "cmpe281"
var mongodb_collection = "starbucksTeas"


func inventoryHandlerBreads(rw http.ResponseWriter, req *http.Request) {


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

func cartHandlerBreads(rw http.ResponseWriter, req *http.Request) {


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

func starbucksAddToCartHandlerBreads(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }


     

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




func starbucksProcessOrdersHandlerBreads(rw http.ResponseWriter, req *http.Request) {


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




func likeHandlerBreads(rw http.ResponseWriter, req *http.Request) {


        if origin := req.Header.Get("Origin"); origin != "" {
        rw.Header().Set("Access-Control-Allow-Origin", origin)
        rw.Header().Set("Content-Type", "application/json")
        rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
        rw.Header().Set("Access-Control-Allow-Headers",
            "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
    }



        fmt.Println(req.FormValue("item"))

        fmt.Println(req.FormValue("likes"))

        //converting likes - string to int
        likes,err := strconv.Atoi(req.FormValue("likes"))
        if err != nil {
            log.Fatal(err)
      } 



    

    session, err := mgo.Dial(mongodb_server)
        if err != nil {
                panic(err)
        }
        defer session.Close()
        session.SetMode(mgo.Monotonic, true)
        c := session.DB(mongodb_database).C(mongodb_collection)



      query := bson.M{"item" : req.FormValue("item")}
      change := bson.M{"$set": bson.M{ "likes" : likes+1 }}

      err = c.Update(query, change)
      if err != nil {
            log.Fatal(err)
      } 
}



func main() {


http.HandleFunc("/inventoryBreads", inventoryHandlerBreads)
http.HandleFunc("/cartItemsBreads", cartHandlerBreads)
http.HandleFunc("/addToCartBreads", starbucksAddToCartHandlerBreads)
http.HandleFunc("/processOrdersBreads", starbucksProcessOrdersHandlerBreads)
http.HandleFunc("/likeIncreaseBreads", likeHandlerBreads)
http.ListenAndServe(":5000", nil)
}


