package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Kiteretzu/golang-projects/03_go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is running on port http://localhost:8080/")
	log.Fatal(http.ListenAndServe("localhost:8080", r))
}
