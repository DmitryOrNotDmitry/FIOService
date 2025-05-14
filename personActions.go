package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var people = []Person{}

func GetPeople(c *gin.Context) {
	c.JSON(http.StatusOK, people)
}

func AddPerson(c *gin.Context) {
	var newPerson Person
	if err := c.ShouldBindJSON(&newPerson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refinePerson(&newPerson)

	people = append(people, newPerson)
	c.JSON(http.StatusCreated, newPerson)
}

func refinePerson(p *Person) {
	p.Age = GetAgeByName(p.Name)
	p.Gender = GetGenderByName(p.Name)
	p.Nation = GetNationByName(p.Name)
}

func UpdatePerson(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var updated Person
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, p := range people {
		if p.Id == id {
			refinePerson(&updated)

			updated.Id = id
			people[i] = updated
			c.JSON(http.StatusOK, updated)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}

func DeletePerson(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	for i, p := range people {
		if p.Id == id {
			people = append(people[:i], people[i+1:]...)
			c.Status(http.StatusNoContent)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
}
