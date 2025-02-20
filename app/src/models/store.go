package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Store represents a store with its details and products
type Store struct {
	StoreID  uuid.UUID `json:"store_id"`
	Name     string    `json:"name"`
	Location string    `json:"location"`
	Products []Product `json:"products"`
}

// MarshalBinary converts the Store struct to a binary format
func (i *Store) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
