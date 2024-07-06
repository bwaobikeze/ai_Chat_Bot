package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ubsss/pp-uu444/sever/cache"
	"github.com/Ubsss/pp-uu444/sever/db"
	"github.com/Ubsss/pp-uu444/sever/requester"
	"github.com/Ubsss/pp-uu444/sever/serverController"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var Router *mux.Router

// get ascii code of hashed uids
// hashUID("uu444") => 390
func hashUID(uid string) int {
	uidLen := len(uid)
	asciCode := 0
	for i := 0; i < uidLen; i++ {
		asciCode += int(uid[i])
	}
	fmt.Println(uidLen)
	return asciCode
}

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading env: %v\n", err.Error())
		return
	}

	url := os.Getenv("CONNECTION_URL")
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	httpRequester := requester.NewHTTPRequester()
	db := db.NewDB(url, username, password)
	cas := cache.NewCache(&httpRequester)
	con := serverController.NewController(&db, &cas)

	Router = mux.NewRouter()
	Router.HandleFunc("/", con.HandleQuery).Methods("POST")
	Router.NotFoundHandler = http.HandlerFunc(con.HandleNotFound)

	// server config
	server := http.Server{
		Addr:         ":8082",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      Router,
	}
	fmt.Println("starting server on port: 8082")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("unable to start server: %v\n", err.Error())
	}
}
