package main

import (
	userService "crud-golang/services"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/users", userService.CreateUser).Methods(http.MethodPost)
	router.HandleFunc("/users", userService.GetUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userService.GetUserByID).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userService.UpdateUser).Methods(http.MethodPut)
	router.HandleFunc("/users/{id}", userService.DeleteUser).Methods(http.MethodDelete)

	fmt.Print("Listenning on 5000")
	log.Fatal(http.ListenAndServe(":5000", router))
}
