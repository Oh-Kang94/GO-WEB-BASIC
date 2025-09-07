package main

import (
	"fmt"
	"net/http"
	"web-basic/src"
	types "web-basic/src/types"
)

func main() {
	s := src.NewServer()
	s.HandleFunc("GET", "/",
		func(c *types.Context) { fmt.Fprintf(c.ResponseWriter, "Welcome!\n") },
	)

	s.HandleFunc("GET", "/about", func(c *types.Context) {
		fmt.Fprintf(c.ResponseWriter, "About Page!\n")
	})

	s.HandleFunc("GET", "/users/:id", func(c *types.Context) {
		if c.Params["id"] == "" {
			panic("id Cannot Be Blank")
		}
		fmt.Fprintf(c.ResponseWriter, "User ID: %v\n", c.Params["id"])
	})

	s.HandleFunc("GET", "/users/:user_id/address/:address_id", func(c *types.Context) {
		fmt.Fprintf(
			c.ResponseWriter,
			"Get User Id: %v\nGet Address: %v\n",
			c.Params["user_id"],
			c.Params["address_id"],
		)
	})

	s.HandleFunc("POST", "/users/:user_id", func(c *types.Context) {
		fmt.Fprintf(c.ResponseWriter, "CREATE for User: %v\n", c.Params["user_id"])
	})

	s.HandleFunc("GET", "/index.html", func(c *types.Context) {
		http.NotFound(c.ResponseWriter, c.Request)
	})

	s.Run(":8000")
}

/*
	❯ curl -d {"name"="오강현"} http://localhost:8000/users/okh1994
	CREATE for User: okh1994
	❯ curl -d name=오강현 http://localhost:8000/users/okh1994
	CREATE for User: okh1994
*/
