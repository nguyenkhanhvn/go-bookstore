package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nguyenkhanhvn/go-bookstore/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	// http.Handle("/", r)
	fmt.Println("Listen and serve port 9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
