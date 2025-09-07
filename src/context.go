package src

import "net/http"

type Context struct {
	Params map[string]any

	ResponseWriter http.ResponseWriter

	Request *http.Request
}

type HandleFunc func(*Context)
