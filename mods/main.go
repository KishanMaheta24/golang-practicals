package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func greeter() {
	fmt.Println("Golang")
}
func main() {
	fmt.Println("Hello mod")
	greeter()
	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>hey there</h1>"))
}
