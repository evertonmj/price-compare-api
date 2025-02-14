package models

import (
	"time"
	"encoding/json"
	"github.com/google/uuid"
)


type Price struct {
	ProductID uuid.UUID `json:product_id`
	StoreID   uuid.UUID `json:store_id`
	Price     float64   `json:price`
	CreatedAt time.Time `json:created_at`
}

func (i *Price) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
}