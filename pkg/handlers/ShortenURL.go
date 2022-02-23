package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/mocks"
	"github.com/washbin/url-shortener/pkg/models"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var data models.Payload

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid payload",
		})
		return
	}
	if !slugRe.MatchString(data.Slug) {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Invalid slug",
		})
		return
	}
	if _, ok := mocks.SiteList[data.Slug]; ok {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusBadRequest),
			Message: "Slug already exists",
		})
		return
	}

	mocks.SiteList[data.Slug] = data.URL

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(models.ShortURL{Short: r.Host + "/" + data.Slug})
	return
}
