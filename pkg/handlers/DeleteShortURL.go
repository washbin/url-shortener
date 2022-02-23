package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/mocks"
	"github.com/washbin/url-shortener/pkg/models"
)

func DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	// remove the initial / by r.URL.Path[1:]
	short_url := r.URL.Path[1:]
	if _, ok := mocks.SiteList[short_url]; ok {
		delete(mocks.SiteList, short_url)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(models.Response{
			Status:  http.StatusText(http.StatusOK),
			Message: "Successfully deleted short url",
		})
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(models.Response{
		Status:  http.StatusText(http.StatusBadRequest),
		Message: "Short url not found",
	})
	return
}
