package models

import (
    "encoding/json"
    "time"

    "github.com/google/uuid"
)

type Product struct {
    ProductID      uuid.UUID `json:"product_id"`
    Name           string    `json:"name"`
    Description    string    `json:"description"`
    CurrentPrice   string   `json:"current_price"`
    HistoricPrices []string   `json:"historic_prices"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (i *Product) MarshalBinary() ([]byte, error) {
    return json.Marshal(i)
}

func (i *Product) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, i)
}