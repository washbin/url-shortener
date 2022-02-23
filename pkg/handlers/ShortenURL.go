package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/models"
)

func (h handler) ShortenURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data models.Payload

	// Check incoming request body
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid payload",
		})
		return
	}
	// Verify slug is valid
	if !slugRe.MatchString(data.Slug) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid slug",
		})
		return
	}
	// Check there is already an existing entry for given slug
	if err := h.DB.Get(ctx, data.Slug).Err(); err == nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Slug already exists",
		})
		return
	}
	// Add a mapping with the slug and given url
	if err := h.DB.Set(ctx, data.Slug, data.URL, 0).Err(); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusInternalServerError),
			Message: "Internal server error",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ShortURL{Short: r.Host + "/" + data.Slug})
	return
}
