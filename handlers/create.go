package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"test-api/model"
	"test-api/storage"
	"test-api/utils"
)

type CreateHandler struct {
	Store storage.Store
	UUID  utils.UUIDGenerator
}

func NewCreateHandler(store storage.Store, uuid utils.UUIDGenerator) *CreateHandler {
	return &CreateHandler{Store: store, UUID: uuid}
}

func (handler *CreateHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	customer := handler.getFormValues(r)
	if customer == nil || customer.Name == "" {
		http.Error(w, "Failed to process form values", http.StatusInternalServerError)
		return
	}

	customer.ID = handler.UUID.New()

	if err := handler.Store.Save(*customer); err != nil {
		http.Error(w, fmt.Sprintf("Error saving customer: %s", err), http.StatusInternalServerError)
		return
	}

	handler.displayFormResponse(w)
}

func (handler *CreateHandler) getFormValues(r *http.Request) *model.Customer {
	customer := &model.Customer{
		Name:        r.FormValue("name"),
		Age:         r.FormValue("age"),
		Nationality: r.FormValue("nationality"),
		Investment:  r.FormValue("investment"),
		Fund:        model.Fund{Equities: r.FormValue("fund1")},
	}

	return customer
}

var RootPath = "./"

func (handler *CreateHandler) displayFormResponse(w http.ResponseWriter) {
	t, err := template.ParseFiles(filepath.Join(RootPath, "static/form-response.html"))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing template: %s", err), http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}
}
