package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/Ubsss/pp-uu444/client/cache"
	"github.com/Ubsss/pp-uu444/client/controller"
	"github.com/Ubsss/pp-uu444/client/query"
	"github.com/Ubsss/pp-uu444/client/requester"
	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	httpRequester := requester.NewHTTPRequester()
	router = mux.NewRouter()
	que := query.NewQuery()
	cas := cache.NewCache(&httpRequester)
	con := controller.NewController(&que, &cas)
	router.HandleFunc("/", con.HandleRequest).Methods("POST")
	router.NotFoundHandler = http.HandlerFunc(con.HandleNotFound)

	// server config
	server := http.Server{
		Addr:         ":8081",
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      router,
	}
	fmt.Println("starting client on port: 8081")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("unable to start client: %v\n", err.Error())
	}
}
