package main

import (
	"fioservice/db"
	"fioservice/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var personsdb *db.PersonsDB = db.CreatePersonsDB()

func GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, personsdb.Get())
}

func AddPerson(c *gin.Context) {
	var newPerson entity.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refinePerson(&newPerson)

	personsdb.Add(&newPerson)
	c.JSON(http.StatusCreated, newPerson)
}

func refinePerson(p *entity.Person) {
	p.Age = GetAgeByName(p.Name)
	p.Gender = GetGenderByName(p.Name)
	p.Nation = GetNationByName(p.Name)
}

func UpdatePerson(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var updated entity.Person
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated.Id = id
	personsdb.Update(&updated)
	c.JSON(http.StatusOK, updated)

	//c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}

func DeletePerson(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	personsdb.Delete(id)
	c.Status(http.StatusNoContent)

	//c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}
