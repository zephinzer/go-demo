package logger

import (
	"log"
	"net/http"
)

func ServerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		secure := ""
		if r.TLS != nil {
			secure = "s"
		}
		log.Printf("@[%s] > %s %s http%s://%s%s", r.RemoteAddr, r.Proto, r.Method, secure, r.Host, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}