package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	var PORT string = ":8080"

	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port: %s", PORT)
	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}

	if r.Method != "GET" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello World!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	email := r.FormValue("email")
	tel := r.FormValue("tel")

	fmt.Fprintf(w, "Form name=%s email=%s tel=%s", name, email, tel)
}
