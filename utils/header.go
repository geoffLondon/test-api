package utils

import "net/http"

func SetHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "Application/json")
}
