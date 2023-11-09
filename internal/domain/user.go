package domain

type User struct {
	ID           string `json:"id" bson:"_id"`
	FullName     string `json:"full_name" bson:"full_name"`
	Email        string `json:"email" bson:"email"`
	Password     string `json:"password,omitempty" bson:"-"`
	PasswordHash string `json:"-" bson:"password_hash"`
	Address      string `json:"address" bson:"address"`
	City         string `json:"city" bson:"city"`
	State        string `json:"state" bson:"state"`
	ZipCode      string `json:"zip_code" bson:"zip_code"`
	Country      string `json:"country" bson:"country"`
}
