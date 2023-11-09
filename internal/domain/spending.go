package domain

import (
	"time"
)

type PaymentMethod string
type Category string

const (
	PaymentMethodPix  PaymentMethod = "pix"
	PaymentMethodCard PaymentMethod = "cartão"
	PaymentMethodCash PaymentMethod = "dinheiro"

	CategoryRestaurant    Category = "restaurante"
	CategoryFuel          Category = "combustível"
	CategoryLeisure       Category = "lazer"
	CategoryMiscellaneous Category = "diversos"
	CategoryClothing      Category = "vestuário"
	CategoryTravel        Category = "viagem"
	CategoryCinema        Category = "cinema"
)

type Spending struct {
	ID            string        `bson:"_id,omitempty" json:"id,omitempty"`
	Author        string        `bson:"author" json:"author"`
	Title         string        `bson:"title" json:"title"`
	Date          time.Time     `bson:"date" json:"date"`
	Value         float64       `bson:"value" json:"value"`
	PaymentMethod PaymentMethod `bson:"payment_method" json:"payment_method"`
	Category      Category      `bson:"category" json:"category"`
	Description   string        `bson:"description" json:"description"`
	CreatedAt     time.Time     `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time     `bson:"updated_at" json:"updated_at"`
	LastEditedBy  string        `bson:"last_edited_by,omitempty" json:"last_edited_by,omitempty"`
}

func (p PaymentMethod) IsValid() bool {
	switch p {
	case PaymentMethodPix, PaymentMethodCard, PaymentMethodCash:
		return true
	default:
		return false
	}
}

func (c Category) IsValid() bool {
	switch c {
	case CategoryRestaurant, CategoryFuel, CategoryLeisure, CategoryMiscellaneous, CategoryClothing, CategoryTravel, CategoryCinema:
		return true
	default:
		return false
	}
}
