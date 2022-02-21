package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Logging logs all requests with its ath and the time it took to process
func Logging() MiddleWare {
	return func(hf http.HandlerFunc) http.HandlerFunc {
		return func(rw http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.Method, r.RequestURI, r.RemoteAddr, time.Since(start))
			}()
			hf(rw, r)
		}
	}
}
