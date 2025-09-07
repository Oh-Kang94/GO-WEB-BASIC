package types

import (
	"net/http"
)

type Context struct {
	Params map[string]any

	ResponseWriter http.ResponseWriter

	Request *http.Request
}

type HandleFunc func(*Context)

type Router struct {
	Handlers map[string]map[string]HandleFunc
}

type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}
