package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/models"
)

func (h handler) DeleteShortURL(w http.ResponseWriter, r *http.Request) {
	// remove the initial / by r.URL.Path[1:]
	slug := r.URL.Path[1:]
	if err := h.DB.Get(ctx, slug).Err(); err == nil {
		h.DB.Del(ctx, slug)

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
