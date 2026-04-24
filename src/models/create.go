package models

type Login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=3"`
}

type NewTransaction struct {
	Type        string  `json:"type" validate:"required,oneof=incoming outgoing" bson:"type"`
	Value       float64 `json:"value" validate:"required,gt=0" bson:"value"`
	Description string  `json:"description" validate:"required" bson:"description"`
}

type NewUser struct {
	Name     string `json:"name" bson:"name" validate:"required"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=3"`
}
