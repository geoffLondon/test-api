package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"test-api/storage"
)

type DeleteByIdHandler struct {
	Store *storage.JSONStore
}

func NewDeleteByIdHandler(store *storage.JSONStore) *DeleteByIdHandler {
	return &DeleteByIdHandler{Store: store}
}

func (handler DeleteByIdHandler) DeleteByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := handler.Store.DeleteByID(id); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}
}
