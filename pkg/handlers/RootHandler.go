package handlers

import (
	"net/http"
)

func RootHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	switch {
	// POST /
	case r.Method == http.MethodPost && r.URL.Path == "/":
		ShortenURL(rw, r)
		return
	// GET /{slug}
	case r.Method == http.MethodGet && redirectURLre.MatchString(r.URL.Path):
		RedirectURL(rw, r)
		return
	// DELETE /{slug}
	case r.Method == http.MethodDelete && redirectURLre.MatchString(r.URL.Path):
		DeleteShortURL(rw, r)
		return
	default:
		NotFound(rw, r)
		return
	}
}
