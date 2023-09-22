package main

import (
	"api-with-interfaces/handlers"
	"log"
	"net/http"

	"api-with-interfaces/router"
	"api-with-interfaces/storage"
	"api-with-interfaces/utils"
)

func main() {
	store := storage.NewJSONStore("storage/customer.json")
	uuidGenerator := utils.UUIDGeneratorImpl{}
	createHandler := handlers.NewCreateHandler(store, &uuidGenerator)
	readHandler := handlers.NewReadHandler(store)
	readByIdHandler := handlers.NewReadByIDHandler(store)
	deleteHandler := handlers.NewDeleteHandler(store)
	deleteByIdHandler := handlers.NewDeleteByIdHandler(store)

	muxRouter := router.MuxRouter(createHandler, readHandler, readByIdHandler, deleteHandler, deleteByIdHandler)

	log.Println("server running...")
	if err := http.ListenAndServe(":8080", muxRouter); err != nil {
		log.Fatal(err)
	}
}
