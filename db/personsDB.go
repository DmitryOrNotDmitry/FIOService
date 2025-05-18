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
	logger.Log.Debugf("Создание нового человека: %+v", p)

	err := db.DB.Create(p).Error
	if err != nil {
		logger.Log.Infof("Ошибка при создании человека: %v", err)
		return err
	}

	logger.Log.Infof("Человек успешно создан: ID=%d", p.Id)
	return nil
}

func (db *PersonsDB) Get(filter *entity.PersonFilter) ([]entity.Person, error) {
	logger.Log.Debug("Запрос на поиск людей с фильтрами")

	var persons []entity.Person
	query := db.DB.Model(&entity.Person{})

	if filter.Name != nil {
		query = query.Where("name ILIKE ?", "%"+*filter.Name+"%")
	}
	if filter.Surname != nil {
		query = query.Where("surname ILIKE ?", "%"+*filter.Surname+"%")
	}
	if filter.Patronymic != nil {
		query = query.Where("patronymic ILIKE ?", "%"+*filter.Patronymic+"%")
	}
	if filter.Gender != nil {
		query = query.Where("gender = ?", *filter.Gender)
	}
	if filter.Nation != nil {
		query = query.Where("nation = ?", *filter.Nation)
	}
	if filter.MinAge != nil {
		query = query.Where("age >= ?", *filter.MinAge)
	}
	if filter.MaxAge != nil {
		query = query.Where("age <= ?", *filter.MaxAge)
	}
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	tx := query.Find(&persons)
	if tx.Error != nil {
		logger.Log.Infof("Ошибка при поиске людей: %v", tx.Error)
		return nil, tx.Error
	}

	logger.Log.Infof("Найдено %d человек", len(persons))
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
