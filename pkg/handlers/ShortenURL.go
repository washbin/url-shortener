package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/models"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data models.Payload
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid request body"}`))
		return
	}
	if !slugRe.MatchString(data.Slug) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Invalid slug"}`))
		return
	}
	if _, ok := siteList[data.Slug]; ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "Slug already exists"}`))
		return
	}
	siteList[data.Slug] = data.URL

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"short_url": "` + r.Host + "/" + data.Slug + `"}`))
	return
}
