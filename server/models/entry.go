package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//So we are Defining our Struct  to be able to access the All the Ingredients
//Each Struct takes in a datatype(That Golang will Understand ) and Will Omit in particular format(in this case JSON)
//First Entry is to be able to Create Object Key in the Mongodb Database
type Entry struct {
	ID          primitive.ObjectID `bson:"id"`
	Dish        *string            `json:"dish"`
	Fat         *float64           `json:"fat"`
	Ingredients *string            `json:"ingredients"`
	Calories    *string            `json:"calories"`
}
