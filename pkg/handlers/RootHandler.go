package handlers

import (
	"net/http"
)

func (h handler) RootHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	switch {
	// POST /
	case r.Method == http.MethodPost && r.URL.Path == "/":
		h.ShortenURL(rw, r)
		return
	// GET /{slug}
	case r.Method == http.MethodGet && redirectURLre.MatchString(r.URL.Path):
		h.RedirectURL(rw, r)
		return
	// DELETE /{slug}
	case r.Method == http.MethodDelete && redirectURLre.MatchString(r.URL.Path):
		h.DeleteShortURL(rw, r)
		return
	default:
		NotFound(rw, r)
		return
	}
}
