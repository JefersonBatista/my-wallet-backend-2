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
	ID          bson.ObjectID `bson:"_id,omitempty" json:"_id"`
	UserID      bson.ObjectID `bson:"userId,omitempty" json:"userId"`
	Type        string        `bson:"type" json:"type"`
	Value       float64       `bson:"value" json:"value"`
	Description string        `bson:"description" json:"description"`
	Timestamp   uint          `bson:"timestamp" json:"timestamp"`
}

type TransactionList struct {
	User string        `json:"user"`
	List []Transaction `json:"list"`
}
