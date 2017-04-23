package middlewares

import (
    "fmt"
    "gopkg.in/gin-gonic/gin.v1"
     "../db"
)

func Connect(c *gin.Context) {
    s := db.Session.Clone()
    defer s.Close()

    c.Set("store", s.DB(db.Mongo.Database))
    c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {

        c.Writer.Header().Set("Content-Type", "application/json")
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Methods","GET,PUT,POST,DELETE,OPTIONS")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Origin, X-Requested-With, Accept, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
        if c.Request.Method == "OPTIONS" {
            fmt.Println(c.Request.Method)
            c.AbortWithStatus(200)
            return
        }

        c.Next()
    }
}