package middleware

import (
	"encoding/json"
	"fmt"
	"web-basic/src/types"
)

func ParseFormHandler(next types.HandleFunc) types.HandleFunc {
	return func(ctx *types.Context) {
		ctx.Request.ParseForm()
		fmt.Println(ctx.Request.PostForm)
		for k, v := range ctx.Request.PostForm {
			if len(v) > 0 {
				ctx.Params[k] = v[0]
			}
		}
		next(ctx)
	}
}

func ParseJsonHandler(next types.HandleFunc) types.HandleFunc {
	return func(ctx *types.Context) {
		var m map[string]any
		if json.NewDecoder(ctx.Request.Body).Decode(&m); len(m) > 0 {
			for k, v := range m {
				ctx.Params[k] = v
			}
		}
		next(ctx)
	}
}
