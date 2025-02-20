package models

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// Product represents a product with its details and prices
type Product struct {
	ProductID      uuid.UUID `json:"product_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	CurrentPrice   string    `json:"current_price"`
	HistoricPrices []string  `json:"historic_prices"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// MarshalBinary converts the Product struct to a binary format
func (i *Product) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}

// UnmarshalBinary converts binary data to the Product struct
func (i *Product) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, i)
}
