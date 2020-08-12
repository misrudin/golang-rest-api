package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/kelas/{perPage}/{page}", getKelas).Methods("GET")
	http.Handle("/", router)
	fmt.Println("Connected to port : 1111")
	log.Fatal(http.ListenAndServe(":1111", router))
}
