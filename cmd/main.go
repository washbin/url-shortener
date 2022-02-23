package main

import (
	"log"
	"net/http"
	"os"

	"github.com/washbin/url-shortener/pkg/db"
	"github.com/washbin/url-shortener/pkg/handlers"
	"github.com/washbin/url-shortener/pkg/middlewares"
)

var RedisAddr = os.Getenv("REDIS_ADDRESS")

// Endpoints
// POST / => create a new short url
// GET /{short_url} => redirect to original url
func main() {
	DB, err := db.Init(RedisAddr)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}
	h := handlers.New(DB)

	mux := http.NewServeMux()
	mux.HandleFunc("/", middlewares.Chain(h.RootHandler, middlewares.Logging()))

	log.Println("Listening on port 8080")
	log.Fatalln(http.ListenAndServe("0.0.0.0:8080", mux))
}
