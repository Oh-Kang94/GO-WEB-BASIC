package src

import (
	"net/http"
	"strings"
	"web-basic/src/types"
)

type Router types.Router

func (r *Router) HandleFunc(method, pattern string, h types.HandleFunc) {
	m, ok := r.Handlers[method]
	if !ok {
		// 등록된 Map이 없으면 생성
		m = make(map[string]types.HandleFunc)
		r.Handlers[method] = m

	}

	m[pattern] = h
}

func (r *Router) Handler() types.HandleFunc {
	return func(ctx *types.Context) {
		for pattern, handler := range r.Handlers[ctx.Request.Method] {
			if ok, params := match(pattern, ctx.Request.URL.Path); ok {
				for k, v := range params {
					ctx.Params[k] = v
				}
				handler(ctx)
				return
			}
		}
		http.NotFound(ctx.ResponseWriter, ctx.Request)
	}

}

func match(pattern, path string) (bool, map[string]string) {
	// 패턴과 Path일치하면 true
	if pattern == path {
		return true, nil
	}

	// Pattern , path를 '/'로 구분
	patterns := strings.Split(pattern, "/")
	paths := strings.Split(path, "/")

	if len(patterns) != len(paths) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range patterns {
		switch {
		case patterns[i] == paths[i]:
		case len(patterns[i]) > 0 && patterns[i][0] == ':':
			params[patterns[i][1:]] = paths[i]
		default:
			return false, nil
		}

	}

	return true, params
}
