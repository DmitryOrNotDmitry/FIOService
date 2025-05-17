package main

import (
	"fioservice/db"
	"fioservice/logger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	logger.Init(logrus.InfoLevel)

	var err error
	err = migrateDB()
	if err != nil {
		logger.Log.Infof("Error occurs with migrate db: %v", err.Error())
		return
	}

	personsdb, err = db.CreatePersonsDB()
	if err != nil {
		logger.Log.Infof("Error occurs with personsdb at Postgres: %v", err.Error())
		return
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/persons", GetPeople)
	r.POST("/persons", AddPerson)
	r.PUT("/persons/:id", UpdatePerson)
	r.DELETE("/persons/:id", DeletePerson)

	logger.Log.Infof("Server runs at port: %v", 8080)
	r.Run(":8080")
}

func migrateDB() error {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:postgres@localhost:5432/personsdb?sslmode=disable")
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
