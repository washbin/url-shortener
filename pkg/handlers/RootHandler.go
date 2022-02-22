package handlers

import (
	"net/http"
	"regexp"
)

var (
	siteList      = make(map[string]string)
	redirectURLre = regexp.MustCompile(`^\/([a-zA-Z0-9]+)$`)
	slugRe        = regexp.MustCompile(`^[a-zA-Z0-9]+$`)
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
	default:
		NotFound(rw, r)
		return
	}
}
