package main

import (
	"fioservice/db"
	"fioservice/logger"
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		return
	}

	logLevel := os.Getenv("APP_LOG_LEVEL")
	switch logLevel {
	case "DEBUG":
		logger.Init(logrus.DebugLevel)
	case "INFO":
		logger.Init(logrus.InfoLevel)
	default:
		logger.Init(logrus.InfoLevel)
	}

	err = migrateDB()
	if err != nil {
		logger.Log.Infof("Ошибка миграции БД: %v", err.Error())
		return
	}

	personsdb, err = db.CreatePersonsDB(getDBOptions())
	if err != nil {
		logger.Log.Infof("Ошибка подключения к БД personsdb: %v", err.Error())
		return
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/persons", GetPeople)
	r.POST("/persons", AddPerson)
	r.PUT("/persons/:id", UpdatePerson)
	r.DELETE("/persons/:id", DeletePerson)

	var port string = os.Getenv("APP_PORT")
	logger.Log.Infof("Сервер запущен на порте: %v", port)
	r.Run(":" + port)
}

func migrateDB() error {
	m, err := migrate.New(
		"file://migrations",
		getDBUrl())
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func getDBOptions() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode)
}

func getDBUrl() string {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	return fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbSSLMode)
}
