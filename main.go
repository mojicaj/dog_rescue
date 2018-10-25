package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/globalsign/mgo"
	"github.com/julienschmidt/httprouter"
	"github.com/mojicaj/dog_rescue/controllers"
	"github.com/mojicaj/dog_rescue/db"
)

var (
	dbServer, database, dbUser, dbPassword, dbURL string
	session                                       *mgo.Session
)

func init() {
	// retrieve database information from env variables
	dbServer = os.Getenv("DB_SERVER")
	database = os.Getenv("DATABASE")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbURL = fmt.Sprintf("mongodb://%s:%s@%s/%s", dbUser, dbPassword, dbServer, database)
}

func main() {
	port := os.Getenv("PORT")
	// open database session
	session, err := mgo.Dial(dbURL)
	if err != nil {
		log.Panic("failed to connect to database: ", err)
	}
	defer session.Close()
	db.Collection = session.DB(database).C("dogs")

	router := httprouter.New()

	// route requests to respective handlers
	router.POST("/api/dog", controllers.CreateDogHandler)
	router.GET("/api/dog/:name", controllers.GetDogHandler)
	router.GET("/api/dog", controllers.GetDogHandler)
	router.PUT("/api/dog", controllers.UpdateDogHandler)
	router.DELETE("/api/dog/:name", controllers.RemoveDogHandler)

	log.Println("Server is listening on port ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
