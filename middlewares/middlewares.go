package middlewares

import (
    "gopkg.in/gin-gonic/gin.v1"
     "../db"
)

func Connect(c *gin.Context) {
    s := db.Session.Clone()
    defer s.Close()

    c.Set("store", s.DB(db.Mongo.Database))
    c.Next()
}