package main

import (
    "fmt"
	"log"
	
	"net/http"

    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
	
    "goSensitive/sensitivefilter/IndexController"
    "goSensitive/sensitivefilter/model"
    "goSensitive/sensitivefilter/WordFilter"
)


func main() {
    session, err := mgo.Dial("127.0.0.1:27017")
    if err != nil {
        panic(err)
    }
    defer session.Close()

    // Optional. Switch the session to a monotonic behavior.
    session.SetMode(mgo.Monotonic, true)

    var list []model.WordStruct
    
    session.DB("test").C("sensitive").Find(bson.M{}).All(&list)

    set := make([]string, 10)
    for index,value := range list {
        fmt.Println(index)
        fmt.Println(value)
        set = append(set,value.Word)
    }
    WordFilter.LoadSensitiveWord(set)

    fmt.Println("init http server ....")
    router := IndexController.InitRoute(session);

    log.Fatal(http.ListenAndServe(":8080", router))
}
