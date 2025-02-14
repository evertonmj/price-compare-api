package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Store struct {
	StoreID  uuid.UUID `json:store_id`
	Name     string    `json:name`
	Location string    `json:location`
	Products []Product `json:products`
}

func (i *Store) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}