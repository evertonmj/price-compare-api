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
    CurrentPrice   float64   `json:"current_price"`
    HistoricPrices []float64   `json:"historic_prices"`
    CreatedAt      time.Time `json:"created_at"`
    UpdatedAt      time.Time `json:"updated_at"`
}

func (i *Product) MarshalBinary() ([]byte, error) {
    return json.Marshal(i)
}

func (i *Product) UnmarshalBinary(data []byte) error {
    return json.Unmarshal(data, i)
}