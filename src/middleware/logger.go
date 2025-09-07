package middleware

import (
	"log"
	"time"
	"web-basic/src/types"
)

func LogHandler(next types.HandleFunc) types.HandleFunc {
	return func(ctx *types.Context) {
		t := time.Now()

		next(ctx)

		log.Printf("[%s] %q %v\n",
			ctx.Request.Method,
			ctx.Request.URL.String(),
			time.Since(t))
	}
}
