package main

import (
	"fmt"
	"log"
	"net/http"
)

const PORT string = "8000"

func helloHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "hello!")
}

func formHandler(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")

	name := req.FormValue("name")
	address := req.FormValue("address")

	fmt.Fprintf(w, "Name: %s\nAddress: %s\n", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir(("./static")))
	http.Handle("/", fileServer)
	http.HandleFunc("/formresult", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server on port %s\n", PORT)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error gay")
	}
}
