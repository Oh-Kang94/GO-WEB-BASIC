package src

import (
	"net/http"
	"web-basic/src/middleware"
	"web-basic/src/types"
)

type Server struct {
	*Router
	middlewares  []middleware.Middleware
	startHandler types.HandleFunc
}

func NewServer() *Server {
	r := &Router{Handlers: make(map[string]map[string]types.HandleFunc)}
	s := &Server{Router: r}
	s.middlewares = []middleware.Middleware{
		middleware.LogHandler,
		middleware.RecoverHandler,
		middleware.StaticHandler,
		middleware.ParseFormHandler,
		middleware.ParseJsonHandler,
	}
	return s
}

func (s *Server) Run(addr string) {
	s.startHandler = s.Router.Handler()

	for i := len(s.middlewares) - 1; i >= 0; i-- {
		s.startHandler = s.middlewares[i](s.startHandler)
	}

	if err := http.ListenAndServe(addr, s); err != nil {
		panic(err)
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := &types.Context{
		Params:         make(map[string]any),
		ResponseWriter: w,
		Request:        r,
	}

	for k, v := range r.URL.Query() {
		c.Params[k] = v[0]
	}
	s.startHandler(c)
}
