package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Chain(hf http.HandlerFunc, middle ...Middleware) http.HandlerFunc {
	for _, m := range middle {
		hf = m(hf)
	}
	return hf
}