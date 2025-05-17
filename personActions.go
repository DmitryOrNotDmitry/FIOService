package main

import (
	"fioservice/db"
	"fioservice/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var personsdb *db.PersonsDB

func GetPeople(c *gin.Context) {
	peoples, err := personsdb.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, peoples)
	}
}

func AddPerson(c *gin.Context) {
	var newPerson entity.Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refinePerson(&newPerson)

	err := personsdb.Add(&newPerson)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, newPerson)
	}
}

func refinePerson(p *entity.Person) {
	p.Age, _ = GetAgeByName(p.Name)
	p.Gender, _ = GetGenderByName(p.Name)
	p.Nation, _ = GetNationByName(p.Name)
}

func UpdatePerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var updated entity.Person
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated.Id = id
	err = personsdb.Update(&updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, updated)
	}
}

func DeletePerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = personsdb.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}
