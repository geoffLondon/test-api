package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"api-with-interfaces/storage"
	"api-with-interfaces/utils"
)

type ReadHandler struct {
	Store storage.Store
}

func NewReadHandler(store storage.Store) *ReadHandler {
	return &ReadHandler{Store: store}
}

func (handler *ReadHandler) Read(w http.ResponseWriter, _ *http.Request) {
	utils.SetHeader(w)
	customers, err := handler.Store.ReadAll()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(customers); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}
}
