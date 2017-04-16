package db

import (
    "fmt"
    "gopkg.in/mgo.v2"
)

var (
    Session *mgo.Session
    Mongo *mgo.DialInfo
)

const (
    MongoDBUrl = "localhost:27017"
)

func Connect() {
    uri := MongoDBUrl
    mongo, err := mgo.ParseURL(uri)
    s, err := mgo.Dial(uri)
    if err != nil {
        fmt.Printf("Can't connect to mongo, go error %v\n", err)
        panic(err.Error())
    }
    s.SetSafe(&mgo.Safe{})
    fmt.Println("Connected to", uri)
    Session = s
    Mongo = mongo
}