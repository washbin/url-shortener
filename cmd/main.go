package main

import (
	"log"
	"net/http"

	"github.com/washbin/url-shortener/pkg/handlers"
	"github.com/washbin/url-shortener/pkg/middlewares"
)

// Endpoints
// POST / => create a new short url
// GET /{short_url} => redirect to original url
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewares.Chain(handlers.RootHandler, middlewares.Logging()))

	log.Println("Listening on port 8080")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", mux))
}
