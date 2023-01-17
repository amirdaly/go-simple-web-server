package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParsForm error", err)
	}
	if r.URL.Path != "/form" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Post form succesfull.\n")
	name := r.FormValue("name")
	family := r.FormValue("family")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Family: %s\n", family)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)
	http.Handle("/", fileServer)

	fmt.Println("Startin web server on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
