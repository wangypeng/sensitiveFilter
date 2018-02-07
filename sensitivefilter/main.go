package main

import (
    "fmt"
	"log"
	
	"net/http"

    "gopkg.in/mgo.v2"
	
	"github.com/julienschmidt/httprouter"

	"goSensitive/sensitivefilter/IndexController"
)


func main() {
    session, err := mgo.Dial("localhost:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    router := httprouter.New()
    router.GET("/", IndexController.Index)
    router.GET("/hello/:name", IndexController.Hello)

    log.Fatal(http.ListenAndServe(":8080", router))
}
