package middleware

import "web-basic/src"

type Middleware func(next src.HandleFunc) src.HandleFunc
