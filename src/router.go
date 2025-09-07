package src

import (
	"net/http"
)

type router struct {
	handlers map[string]map[string]func(http.ResponseWriter, *http.Request) // http.HandlerFunc Mapping
}

type Handler interface {
	ServeHttp(http.ResponseWriter, *http.Request)
}

func (r *router) HandleFunc(method, pattern string, h http.HandlerFunc) {
	m, ok := r.handlers[method]
	if !ok {
		// 등록된 Map이 없으면 생성
		m = make(map[string]func(http.ResponseWriter, *http.Request))
		r.handlers[method] = m

	}

	m[pattern] = h
}

func (r *router) ServeHttp(w http.ResponseWriter, req *http.Request) {
	if m, ok := r.handlers[req.Method]; ok {
		if h, ok := m[req.URL.Path]; ok {
			h(w, req)
			return
		}
	}
	http.NotFound(w, req)
}
