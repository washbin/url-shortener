package handlers

import (
	"net/http"
)

func (h handler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	// remove the initial / by r.URL.Path[1:]
	slug := r.URL.Path[1:]
	// Redirect if there is a mapping for slug
	if url := h.DB.Get(ctx, slug).Val(); url != "" {
		http.Redirect(w, r, url, http.StatusMovedPermanently)
		return
	}

	NotFound(w, r)
	return
}
