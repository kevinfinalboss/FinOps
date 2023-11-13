package domain

import (
	"time"
)

type IncomeMethod string

const (
	IncomeMethodTED    IncomeMethod = "TED"
	IncomeMethodCheque IncomeMethod = "Cheque"
	IncomeMethodCash   IncomeMethod = "Dinheiro"
	IncomeMethodPix    IncomeMethod = "Pix"
	IncomeMethodOther  IncomeMethod = "Outros"
)

type Income struct {
	ID           string       `bson:"_id,omitempty" json:"id,omitempty"`
	Author       string       `bson:"author" json:"author"`
	Title        string       `bson:"title" json:"title"`
	Date         string       `json:"date"`
	Value        float64      `bson:"value" json:"value"`
	IncomeMethod IncomeMethod `bson:"incomeMethod" json:"incomeMethod"`
	Description  string       `bson:"description" json:"description"`
	CreatedAt    time.Time    `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time    `bson:"updated_at" json:"updated_at"`
	LastEditedBy string       `bson:"last_edited_by,omitempty" json:"last_edited_by,omitempty"`
}

func (m IncomeMethod) IsValid() bool {
	switch m {
	case IncomeMethodTED, IncomeMethodCheque, IncomeMethodCash, IncomeMethodPix, IncomeMethodOther:
		return true
	default:
		return false
	}
}
