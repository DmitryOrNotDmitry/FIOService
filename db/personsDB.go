package db

import (
	"fioservice/entity"
	"fioservice/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PersonsDB struct {
	DB *gorm.DB
}

func (db *PersonsDB) Add(p *entity.Person) error {
	logger.Log.Debugf("Добавление нового человека: %+v", p)

	err := db.DB.Create(p).Error
	if err != nil {
		logger.Log.Infof("Ошибка при добавлении человека: %v", err)
		return err
	}

	logger.Log.Infof("Человек успешно добавлен: ID=%d", p.Id)
	return nil
}

func (db *PersonsDB) Get() ([]entity.Person, error) {
	logger.Log.Debug("Запрос на получение всех людей")

	var persons []entity.Person
	tx := db.DB.Find(&persons)
	if tx.Error != nil {
		logger.Log.Infof("Ошибка при получении людей: %v", tx.Error)
		return nil, tx.Error
	}

	logger.Log.Infof("Получено %d человек", len(persons))
	return persons, nil
}

func (db *PersonsDB) Update(updated *entity.Person) error {
	logger.Log.Debugf("Обновление человека: %+v", updated)

	err := db.DB.Model(updated).
		Where("id = ?", updated.Id).
		Updates(updated).
		Error
	if err != nil {
		logger.Log.Infof("Ошибка при обновлении человека ID=%d: %v", updated.Id, err)
		return err
	}

	logger.Log.Infof("Человек ID=%d успешно обновлен", updated.Id)
	return nil
}

func (db *PersonsDB) Delete(id int64) error {
	logger.Log.Debugf("Удаление человека с ID=%d", id)

	err := db.DB.Delete(&entity.Person{Id: id}).Error
	if err != nil {
		logger.Log.Infof("Ошибка при удалении человека с ID=%d: %v", id, err)
		return err
	}

	logger.Log.Infof("Человек ID=%d успешно удален", id)
	return nil
}

func CreatePersonsDB(dbOptions string) (*PersonsDB, error) {
	logger.Log.Debug("Попытка подключения к базе данных")
	DB, err := gorm.Open(postgres.Open(dbOptions), &gorm.Config{})
	if err != nil {
		logger.Log.Infof("Ошибка подключения к БД: %v", err)
		return nil, err
	}

	logger.Log.Info("Подключение к базе данных успешно")
	return &PersonsDB{DB}, nil
}
