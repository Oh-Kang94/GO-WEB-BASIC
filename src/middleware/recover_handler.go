package middleware

import (
	"log"
	"net/http"
	"web-basic/src"
)

func RecoverHandler(next src.HandleFunc) src.HandleFunc {
	return func(ctx *src.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic : %+v", err)
				http.Error(
					ctx.ResponseWriter,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next(ctx)
	}
}
