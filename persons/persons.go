package persons

import (
    "gopkg.in/gin-gonic/gin.v1"
    "gopkg.in/mgo.v2"
    "gopkg.in/mgo.v2/bson"
    "net/http"
    "../models"
)

func Create(c *gin.Context) {
    db := c.MustGet("store").(*mgo.Database)

    persons := models.Person{}
    err := c.BindJSON(&persons)
    if err != nil {
        c.Error(err)
        return
    }
    persons.Id = bson.NewObjectId()

    err = db.C(models.CollectionPerson).Insert(persons)
    if err != nil {
        c.Error(err)
    }
    c.JSON(201, persons)
}

func Update(c *gin.Context) {
    db := c.MustGet("store").(*mgo.Database)

    persons := models.Person{}
    err := c.BindJSON(&persons)
    if err != nil {
        c.Error(err)
        return
    }

    query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}

    persons.Id = bson.ObjectIdHex(c.Param("_id"))

    err = db.C(models.CollectionPerson).Update(query, persons)
    if err != nil {
        c.Error(err)
    }

    c.Writer.WriteHeader(http.StatusNoContent)
}

func List(c *gin.Context) {
    db := c.MustGet("store").(*mgo.Database)
    persons := []models.Person{}
    err := db.C(models.CollectionPerson).Find(nil).Sort("-updated_on").All(&persons)
    if err != nil {
        c.Error(err)
    }

    c.JSON(200, persons)
}

func Get(c *gin.Context) {
    db := c.MustGet("store").(*mgo.Database)
    person := models.Person{}
    query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
    err := db.C(models.CollectionPerson).Find(query).One(&person)
    if err != nil {
        c.Error(err)
    }
    c.JSON(200, person)
}

func Delete(c *gin.Context) {
    db := c.MustGet("store").(*mgo.Database)
    query := bson.M{"_id": bson.ObjectIdHex(c.Param("_id"))}
    err := db.C(models.CollectionPerson).Remove(query)
    if err != nil {
        c.Error(err)
    }
    c.Writer.WriteHeader(http.StatusNoContent)
}