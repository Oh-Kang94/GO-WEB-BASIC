package middleware

import "web-basic/src/types"

type Middleware func(next types.HandleFunc) types.HandleFunc
