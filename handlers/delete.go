package handlers

import (
	"fmt"
	"net/http"

	"test-api/storage"
)

type DeleteHandler struct {
	Store *storage.JSONStore
}

func NewDeleteHandler(store *storage.JSONStore) *DeleteHandler {
	return &DeleteHandler{Store: store}
}

func (handler DeleteHandler) Delete(w http.ResponseWriter, _ *http.Request) {
	if err := handler.Store.DeleteAll(); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}
}
