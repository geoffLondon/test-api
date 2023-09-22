package main

import (
	"log"
	"net/http"

	"test-api/handlers"
	"test-api/router"
	"test-api/storage"
	"test-api/utils"
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
