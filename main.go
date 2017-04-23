package main

import (
    "gopkg.in/gin-gonic/gin.v1"
    "./models"
    "./middlewares"
    "./persons"
    "./db"
)

func index (c *gin.Context) {
    content := gin.H{"Hello": "World"}
    c.JSON(200, content)
}

func organization(c *gin.Context) {
    content := models.Organization{"Control Union"}
    c.JSON(200, content)
}

func init() {
    db.Connect()
}

func main() {
    app := gin.Default()
    app.Use(middlewares.Connect)
    app.Use(middlewares.CORSMiddleware())
    app.GET("/", index)
    app.POST("/persons", persons.Create)
    app.GET("/persons", persons.List)
    app.PUT("/persons/:_id", persons.Update)
    app.GET("/persons/:_id", persons.Get)
    app.DELETE("/persons/:_id", persons.Delete)
    app.GET("/organization", organization)
    app.Run(":8001")
}