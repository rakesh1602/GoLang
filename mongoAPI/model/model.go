package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Netflix struct {
	Id      primitive.ObjectID `json:"id,omitempty"`
	Movie   string             `json:"movie,omitempty"`
	Watched bool               `json:"watched,omitempty"`
}
