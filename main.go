package main

import (
	"fmt"
	"net/http"
	"time"
	"web-basic/src"
	mid "web-basic/src/middleware"
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

	s.HandleFunc("GET", "/users/:id", func(ctx *types.Context) {
		if ctx.Params["id"] == "" {
			panic("id Cannot Be Blank")
		}

		u := model.User{Id: ctx.Params["id"].(string)}
		ctx.RenderResponse(u)
	})

	s.HandleFunc("GET", "/users/:user_id/address/:address_id", func(ctx *types.Context) {
		u := model.User{Id: ctx.Params["user_id"].(string), AddressId: ctx.Params["address_id"].(string)}
		ctx.RenderResponse(u)
	})

	s.HandleFunc("POST", "/users/:user_id", func(c *types.Context) {
		fmt.Fprintf(c.ResponseWriter, "CREATE for User: %v\n", c.Params["user_id"])
	})

	s.HandleFunc("GET", "/", func(ctx *types.Context) {
		ctx.RenderTemplate("src/public/index.html", map[string]any{"time": time.Now()})
	})

	// Login
	s.HandleFunc("GET", "/login", func(ctx *types.Context) {
		ctx.RenderTemplate("src/public/login.html", map[string]any{"msg": "로그인이 필요합니다."})
	})

	s.HandleFunc("POST", "/login", func(ctx *types.Context) {
		if CheckLogin(ctx.Params["id"].(string), ctx.Params["pw"].(string)) {
			http.SetCookie(ctx.ResponseWriter, &http.Cookie{
				Name:  "X_AUTH",
				Value: mid.Sign(mid.VerifyMsg),
				Path:  "/",
			})
			ctx.Redirect("/")
		}

		ctx.RenderTemplate("src/public/login.html", map[string]any{"msg": "id 또는 패스워드가 일치하지 않습니다."})
	})

	s.Run(":8000")
}

func CheckLogin(id, pw string) bool {
	const (
		ID = "test"
		PW = "1234"
	)

	return id == ID && pw == PW
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
