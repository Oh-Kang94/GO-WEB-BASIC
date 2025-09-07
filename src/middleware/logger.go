package middleware

import (
	"log"
	"time"
	"web-basic/src"
)

func LogHandler(next src.HandleFunc) src.HandleFunc {
	return func(ctx *src.Context) {
		t := time.Now()

		next(ctx)

		log.Printf("[%s] %q %v\n",
			ctx.Request.Method,
			ctx.Request.URL.String(),
			time.Since(t))
	}
}
