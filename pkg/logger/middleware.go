package logger

import (
	"net/http"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Sugar.Infof("New request on %s from IP %s.", r.URL, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}
