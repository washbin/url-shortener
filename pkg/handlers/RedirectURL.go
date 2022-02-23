package handlers

import (
	"net/http"

	"github.com/washbin/url-shortener/pkg/mocks"
)

func RedirectURL(w http.ResponseWriter, r *http.Request) {
	// remove the initial / by r.URL.Path[1:]
	if val, ok := mocks.SiteList[r.URL.Path[1:]]; ok {
		http.Redirect(w, r, val, http.StatusMovedPermanently)
		return
	}

	NotFound(w, r)
	return
}
