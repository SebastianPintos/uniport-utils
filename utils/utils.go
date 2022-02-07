package utils

import (
	"go.mongodb.org/mongo-driver/bson"
)

func StringToFilter(field string, value string) bson.D {
	if value == "" {
		return bson.D{}
	}
	return bson.D{{field, value}}
}
