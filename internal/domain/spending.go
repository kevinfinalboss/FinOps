package domain

import (
	"time"
)

type Category string
type PaymentMethod string

const (
	PaymentMethodPix  PaymentMethod = "Pix"
	PaymentMethodCard PaymentMethod = "Cartão"
	PaymentMethodCash PaymentMethod = "Dinheiro"
)

const (
	CategoryRestaurant    Category = "restaurante"
	CategoryFuel          Category = "combustível"
	CategoryLeisure       Category = "lazer"
	CategoryMiscellaneous Category = "diversos"
	CategoryClothing      Category = "vestuário"
	CategoryTravel        Category = "viagem"
	CategoryCinema        Category = "cinema"
)

type Spending struct {
	ID            string    `bson:"_id,omitempty" json:"id,omitempty"`
	Author        string    `bson:"author" json:"author"`
	Title         string    `bson:"title" json:"title"`
	Date          string    `bson:"date" json:"date"`
	Value         float64   `bson:"value" json:"value"`
	Category      Category  `bson:"category" json:"category"`
	PaymentMethod string    `bson:"paymentMethod" json:"paymentMethod"`
	Description   string    `bson:"description" json:"description"`
	CreatedAt     time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time `bson:"updated_at" json:"updated_at"`
	LastEditedBy  string    `bson:"last_edited_by,omitempty" json:"last_edited_by,omitempty"`
}

func (c Category) IsValid() bool {
	switch c {
	case CategoryRestaurant, CategoryFuel, CategoryLeisure, CategoryMiscellaneous, CategoryClothing, CategoryTravel, CategoryCinema:
		return true
	default:
		return false
	}
}
