package models

type Session struct {
	Token string `bson:"token"`
}

type User struct {
	Name         string `json:"name" bson:"name"`
	Email        string `json:"email" bson:"email"`
	PasswordHash string `json:"passwordHash" bson:"passwordHash"`
}

type Transaction struct {
	Type        string  `bson:"type"`
	Value       float64 `bson:"value"`
	Description string  `bson:"description"`
	Timestamp   uint    `bson:"timestamp"`
}
