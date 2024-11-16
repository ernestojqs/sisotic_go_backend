package models

import (
	"otic/resolvers/scalars"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupportStatusDetails struct {
	Type       string             `bson:"type" gql:"name=type"`
	Date       scalars.DateTime   `bson:"date" gql:"name=date"`
	RecordUser primitive.ObjectID `bson:"recordUser" gql:"name=recordUser,objectID=true"`
}
