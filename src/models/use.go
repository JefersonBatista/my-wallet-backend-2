package models

import "go.mongodb.org/mongo-driver/v2/bson"

type Session struct {
	ID     bson.ObjectID `bson:"_id,omitempty"`
	UserID bson.ObjectID `bson:"userId,omitempty"`
	Token  string        `bson:"token"`
}

type User struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Name         string        `json:"name" bson:"name"`
	Email        string        `json:"email" bson:"email"`
	PasswordHash string        `json:"passwordHash" bson:"passwordHash"`
}

type Transaction struct {
	ID          bson.ObjectID `bson:"_id,omitempty"`
	UserID      bson.ObjectID `bson:"userId,omitempty"`
	Type        string        `bson:"type"`
	Value       float64       `bson:"value"`
	Description string        `bson:"description"`
	Timestamp   uint          `bson:"timestamp"`
}

type TransactionList struct {
	User string        `json:"user"`
	List []Transaction `json:"list"`
}
