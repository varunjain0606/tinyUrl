package main

import (
	"fmt"
	"log"
	"net/http"
)

func createUrl(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")

	fmt.Println("Endpoint Hit: homePage")
}

func fetchUrl(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	http.HandleFunc("/v1/tinyUrl/url", createUrl)
	http.HandleFunc("/v1/tinyUrl/url/{shortenUrl}", fetchUrl)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
	handleRequests()
}
