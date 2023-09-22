package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func DisplayForm(w http.ResponseWriter, _ *http.Request) {
	t, _ := template.ParseFiles("./static/form.html")
	if err := t.Execute(w, nil); err != nil {
		http.Error(w, fmt.Sprintf("Error: %s", err), http.StatusBadRequest)
		return
	}
}
