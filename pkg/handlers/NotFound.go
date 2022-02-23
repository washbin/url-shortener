package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/washbin/url-shortener/pkg/models"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(models.Response{
		Status:  http.StatusText(http.StatusNotFound),
		Message: "Not found",
	})
	return
}
