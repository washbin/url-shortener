package handlers

import (
	"net/http"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	if val, ok := siteList[r.URL.Path[1:]]; ok {
		http.Redirect(w, r, val, http.StatusMovedPermanently)
		return
	}

	NotFound(w, r)
	return
}