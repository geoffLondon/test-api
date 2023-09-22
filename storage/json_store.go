package storage

import (
	"encoding/json"
	"io"
	"log"
	"os"

	"test-api/model"
)

type Store interface {
	Save(c model.Customer) error
	ReadAll() ([]model.Customer, error)
	ReadByID(id string) (model.Customer, error)
	DeleteAll() error
	DeleteByID(id string) error
}

type JSONStore struct {
	filePath string
}

func NewJSONStore(path string) *JSONStore {
	return &JSONStore{filePath: path}
}

func (js *JSONStore) Save(customer model.Customer) error {
	customers, err := js.ReadAll()
	if err != nil {
		return err
	}
	customers = append(customers, customer)
	bytes, err := json.Marshal(customers)
	if err != nil {
		return err
	}
	if err := os.WriteFile(js.filePath, bytes, 0644); err != nil {
		return err
	}

	return nil
}

func (js *JSONStore) ReadAll() ([]model.Customer, error) {
	file, err := os.Open(js.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if len(bytes) == 0 {
		return []model.Customer{}, nil
	}
	if err != nil {
		return nil, err
	}

	var customers []model.Customer
	if err := json.Unmarshal(bytes, &customers); err != nil {
		return nil, err
	}
	return customers, nil
}

func (js *JSONStore) ReadByID(id string) (model.Customer, error) {
	customers, err := js.ReadAll()
	if err != nil {
		return model.Customer{}, err
	}

	for _, customer := range customers {
		if customer.ID == id {
			return customer, nil
		}
	}

	log.Printf("person not found with id %v", id)

	return model.Customer{}, nil
}

func (js *JSONStore) DeleteAll() error {
	if err := os.WriteFile(js.filePath, []byte{}, 0644); err != nil {
		return err
	}
	return nil
}

func (js *JSONStore) DeleteByID(id string) error {
	customers, err := js.ReadAll()
	if err != nil {
		return err
	}

	var updatedCustomers []model.Customer
	for _, customer := range customers {
		if customer.ID == id {
			continue
		}
		updatedCustomers = append(updatedCustomers, customer)
	}

	bytes, err := json.Marshal(updatedCustomers)
	if err != nil {
		return err
	}

	if err := os.WriteFile(js.filePath, bytes, 0644); err != nil {
		return err
	}

	return nil
}
