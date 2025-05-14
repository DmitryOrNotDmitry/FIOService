package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/persons", GetPeople)
	r.POST("/persons", AddPerson)
	r.PUT("/persons/:id", UpdatePerson)
	r.DELETE("/persons/:id", DeletePerson)

	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}
