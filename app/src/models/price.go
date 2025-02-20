package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Price represents the price of a product in a store
type Price struct {
	ProductID uuid.UUID `json:"product_id"`
	StoreID   uuid.UUID `json:"store_id"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

// MarshalBinary converts the Price struct to a binary format
func (i *Price) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}
