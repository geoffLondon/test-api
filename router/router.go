package router

import (
	"api-with-interfaces/handlers"
	"github.com/gorilla/mux"
)

func MuxRouter(ch *handlers.CreateHandler, r *handlers.ReadHandler, rbid *handlers.ReadByIDHandler, da *handlers.DeleteHandler, dbid *handlers.DeleteByIdHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/", handlers.DisplayForm).Methods("GET")
	router.HandleFunc("/customer/create", ch.CreateCustomer).Methods("POST")
	router.HandleFunc("/customer", r.Read).Methods("GET")
	router.HandleFunc("/customer/{id}", rbid.ReadByID).Methods("GET")
	router.HandleFunc("/customer", da.Delete).Methods("DELETE")
	router.HandleFunc("/customer/{id}", dbid.DeleteByID).Methods("DELETE")

	return router
}
