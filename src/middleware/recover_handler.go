package middleware

import (
	"log"
	"net/http"
	"web-basic/src/types"
)

func RecoverHandler(next types.HandleFunc) types.HandleFunc {
	return func(ctx *types.Context) {
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
