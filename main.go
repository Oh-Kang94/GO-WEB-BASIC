package main

import (
	"fmt"
	"net/http"
	"web-basic/src"
	mid "web-basic/src/middleware"
)

func main() {

	r := &src.Router{
		Handlers: make(map[string]map[string]src.HandleFunc),
	}

	r.HandleFunc("GET", "/", mid.LogHandler(func(c *src.Context) {
		fmt.Println(c.ResponseWriter, "c.ResponseWriterelcome!")
	}))

	r.HandleFunc("GET", "/about", mid.LogHandler(func(c *src.Context) {
		fmt.Println(c.ResponseWriter, "about!")
	}))
	r.HandleFunc("GET", "/users/:id", mid.LogHandler(func(c *src.Context) {
		fmt.Fprintf(c.ResponseWriter, "user id: %v\n", c.Params["id"])
	}))
	r.HandleFunc("GET", "/users/:user_id/address/:address_id", mid.LogHandler(func(c *src.Context) {
		fmt.Fprintf(c.ResponseWriter, "Get User Id: %v\nGet Address: %v\n", c.Params["user_id"], c.Params["address_id"])
	}))

	r.HandleFunc("POST", "/users/:user_id/addresses", mid.LogHandler(func(c *src.Context) {
		fmt.Println(c.ResponseWriter, "CREATE USER : %v", c.Params["user_id"])
	}))

	http.ListenAndServe(":8000", r)
}
