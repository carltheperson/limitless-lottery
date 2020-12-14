package db

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username,omitempty"`
	PasswordHash string             `bson:"passwordHash,omitempty"`
	Balance      int                `bson:"balance,omitempty"`
}

type SessionIdentity struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Username       string             `bson:"username,omitempty"`
	SessionToken   string             `bson:"sessionToken,omitempty"`
	ExpirationDate int64              `bson:"expirationDate,omitempty"`
}
