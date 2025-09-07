package main

import (
	"fmt"
	"time"
	"web-basic/src"
	"web-basic/src/model"
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
		ctx := (*src.Context)(c) // 순환참조 피하기 위해 여기서 형변환
		if ctx.Params["id"] == "" {
			panic("id Cannot Be Blank")
		}

		u := model.User{Id: ctx.Params["id"].(string)}
		ctx.RenderResponse(u)
	})

	s.HandleFunc("GET", "/users/:user_id/address/:address_id", func(c *types.Context) {
		ctx := (*src.Context)(c) // 순환참조 피하기 위해 여기서 형변환
		u := model.User{Id: ctx.Params["user_id"].(string), AddressId: ctx.Params["address_id"].(string)}
		ctx.RenderResponse(u)
	})

	s.HandleFunc("POST", "/users/:user_id", func(c *types.Context) {
		fmt.Fprintf(c.ResponseWriter, "CREATE for User: %v\n", c.Params["user_id"])
	})

	s.HandleFunc("GET", "/", func(c *types.Context) {
		ctx := (*src.Context)(c) // 순환참조 피하기 위해 여기서 형변환
		ctx.RenderTemplate("src/public/index.html", map[string]any{"time": time.Now()})
	})

	s.Run(":8000")
}

/*
	❯ curl -d {"name"="오강현"} http://localhost:8000/users/okh1994
	CREATE for User: okh1994
	❯ curl -d name=오강현 http://localhost:8000/users/okh1994
	CREATE for User: okh1994
*/

/*
	❯ curl http://127.0.0.1:8000/users/okh1994
	{"Id":"okh1994","AddressId":"","Name":""}

	❯ curl -H "Accept: application/xml" http://127.0.0.1:8000/users/okh1994
	<User><Id>okh1994</Id><AddressId></AddressId><Name></Name></User>%
*/
