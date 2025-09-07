package middleware

import (
	"net/http"
	"path"
	"strings"
	"web-basic/src/types"
)

func StaticHandler(next types.HandleFunc) types.HandleFunc {
	var (
		// 정적 리소스 루트를 types/public 으로 고정
		dir       = http.Dir("./types/public")
		indexFile = "index.html"
	)

	return func(ctx *types.Context) {
		if ctx.Request.Method != "GET" && ctx.Request.Method != "HEAD" {
			next(ctx)
			return
		}

		file := ctx.Request.URL.Path

		f, err := dir.Open(file)
		if err != nil {
			next(ctx)
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			next(ctx)
			return
		}

		if fi.IsDir() {
			if !strings.HasSuffix(ctx.Request.URL.Path, "/") {
				http.Redirect(ctx.ResponseWriter, ctx.Request, ctx.Request.URL.Path+"/", http.StatusFound)
				return
			}

			file = path.Join(file, indexFile)

			f, err = dir.Open(file)
			if err != nil {
				next(ctx)
				return
			}
			defer f.Close()

			fi, err = f.Stat()
			if err != nil || fi.IsDir() {
				next(ctx)
				return
			}
		}

		http.ServeContent(ctx.ResponseWriter, ctx.Request, file, fi.ModTime(), f)
	}
}
