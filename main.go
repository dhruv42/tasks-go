package main

import (
	"log"
	"net/http"
	"tasks-go/routes"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	routes.RegisterTasksRoutes(r)
	log.Fatal(http.ListenAndServe("localhost:8081", r))
}

func ErrorCheck(err error) {
	if err != nil {
		panic(err.Error())
	}
}
