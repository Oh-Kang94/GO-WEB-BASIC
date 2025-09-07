package main

import (
	"fmt"
	"net/http"
	"web-basic/src"
)

func main() {

	r := &src.Router{
		Handlers: make(map[string]map[string]http.HandlerFunc),
	}

	r.HandleFunc("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Welcome!")
	})

	r.HandleFunc("GET", "/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "about!")
	})
	r.HandleFunc("GET", "/users/:id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "user id: \n")
	})
	r.HandleFunc("GET", "/users/:user_id/address/:address_id", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Get User Id: \nGet Address: \n")
	})

	r.HandleFunc("POST", "/users/:user_id/addresses", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "CREATE USER")
	})

	http.ListenAndServe(":8000", r)
}
