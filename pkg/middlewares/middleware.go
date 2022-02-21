package middlewares

import (
	"net/http"
)

type MiddleWare func(http.HandlerFunc) http.HandlerFunc
