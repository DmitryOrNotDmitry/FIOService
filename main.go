package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migrateDB()

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/persons", GetPeople)
	r.POST("/persons", AddPerson)
	r.PUT("/persons/:id", UpdatePerson)
	r.DELETE("/persons/:id", DeletePerson)

	log.Println("Сервер запущен на :8080")
	r.Run(":8080")
}

func migrateDB() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/personsdb?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
