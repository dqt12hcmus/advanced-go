package main

import "net/http"

import "os"

import "fmt"

func main() {
	http.HandleFunc("/info", info)
	http.ListenAndServe(port(), nil)
}

func info(w http.ResponseWriter, r *http.Request) {
	fmt.Println("The /info endpoint is being called")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello K8S discovery & configuration")
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}
