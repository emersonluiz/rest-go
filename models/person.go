package models

import "gopkg.in/mgo.v2/bson"

const (
    CollectionPerson = "persons"
)

type Person struct {
    Id bson.ObjectId `json:"id" bson:"_id"`
    FirstName string `json:"firstName"`
    LastName string `json:"lastName"`
}