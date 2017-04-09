package main

import (
    "gopkg.in/gin-gonic/gin.v1"
)

type Person struct {
    Name string
}

func index (c *gin.Context) {
    content := gin.H{"Hello": "World"}
    c.JSON(200, content)
}

func person(c *gin.Context) {
    content := gin.H{"name": "Emerson Luiz de Castro"}
    c.JSON(200, content)
}

func personList(c *gin.Context) {
    persons := make([]Person, 0)
    persons = append(persons, Person{"Emerson Luiz de Castro"})
    persons = append(persons, Person{"Juca da Silva"})
    c.JSON(200, persons)
}

func main() {
    app := gin.Default()
    app.GET("/", index)
    app.GET("/person", person)
    app.GET("/persons", personList)
    app.Run(":8000")
}