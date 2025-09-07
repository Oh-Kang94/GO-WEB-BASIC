package src

import (
	"net/http"
	"strings"
)

type Router struct {
	Handlers map[string]map[string]HandleFunc
}

type Handler interface {
	ServeHttp(http.ResponseWriter, *http.Request)
}

func (r *Router) HandleFunc(method, pattern string, h HandleFunc) {
	m, ok := r.Handlers[method]
	if !ok {
		// 등록된 Map이 없으면 생성
		m = make(map[string]HandleFunc)
		r.Handlers[method] = m

	}

	m[pattern] = h
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for pattern, handler := range r.Handlers[req.Method] {
		if ok, params := match(pattern, req.URL.Path); ok {
			c := Context{
				Params:         make(map[string]any),
				ResponseWriter: w,
				Request:        req,
			}
			for k, v := range params {
				c.Params[k] = v
			}
			handler(&c)
			return
		}
	}

	http.NotFound(w, req)
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
