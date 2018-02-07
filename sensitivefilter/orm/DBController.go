package indexController

import (
	"fmt"
	"log"

	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Controller struct{
	S *mgo.Session
}

func (controller Controller) Query (w http.ResponseWriter ,r *http.Request){

	session := controller.S.Copy()
	defer session.Close()

	var p []Person

	session.DB("test").C("people").Find(bson.M{}).All(&p)

	for index , value  := range p {
		fmt.Println(index)
		fmt.Println(value)
	}
	log.Println("Failed get all books: ", p)

	fmt.Fprint(w,p);
}

func (controller Controller) Insert (w http.ResponseWriter ,r *http.Request){

	session := controller.S.Copy()
	defer session.Close()

	var p []Person

	session.DB("test").C("people").Find(bson.M{}).All(&p)

	for index , value  := range p {
		fmt.Println(index)
		fmt.Println(value)
	}
	log.Println("Failed get all books: ", p)

	fmt.Fprint(w,p);
}

func (controller Controller) Delete (w http.ResponseWriter ,r *http.Request){

	session := controller.S.Copy()
	defer session.Close()

	var p []Person

	session.DB("test").C("people").Find(bson.M{}).All(&p)

	for index , value  := range p {
		fmt.Println(index)
		fmt.Println(value)
	}
	log.Println("Failed get all books: ", p)

	fmt.Fprint(w,p);
}