package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"test-api/storage"
	"test-api/utils"
)

type ReadByIDHandler struct {
	Store storage.Store
}

func NewReadByIDHandler(store storage.Store) *ReadByIDHandler {
	return &ReadByIDHandler{Store: store}
}

func (handler *ReadByIDHandler) ReadByID(w http.ResponseWriter, r *http.Request) {
	utils.SetHeader(w)
	id := mux.Vars(r)["id"]

	fmt.Println("ID:", id)
	customer, err := handler.Store.ReadByID(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(customer); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusInternalServerError)
		return
	}
}
