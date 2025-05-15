package db

import (
	"fioservice/entity"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonsDB struct {
	DB *gorm.DB
}

func (db *PersonsDB) Add(p *entity.Person) {
	if err := db.DB.Create(p).Error; err != nil {
		panic("Failed to save person: " + err.Error())
	}
}

func (db *PersonsDB) Get() []entity.Person {
	var persons []entity.Person
	tx := db.DB.Find(&persons)

	if tx.Error != nil {
		panic("Failed to save person: " + tx.Error.Error())
	}

	return persons
}

func (db *PersonsDB) Update(updated *entity.Person) {
	db.DB.Model(updated).Where("id = ?", updated.Id).Updates(updated)
}

func (db *PersonsDB) Delete(id int64) {
	db.DB.Delete(&entity.Person{Id: id})
}

func CreatePersonsDB() *PersonsDB {
	dbOptions := "host=localhost user=postgres password=postgres dbname=personsdb port=5432 sslmode=disable"

	DB, err := gorm.Open(postgres.Open(dbOptions), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	fmt.Println("Database connected.")

	return &PersonsDB{DB}
}
