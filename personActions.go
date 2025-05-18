package main

import (
	"fioservice/db"
	"fioservice/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var personsdb *db.PersonsDB

// @Summary Получить список людей с фильтрацией
// @Description Возвращает список людей по заданным фильтрам (имя, фамилия, возраст и др.)
// @Tags persons
// @Accept json
// @Produce json
// @Param name query string false "Имя"
// @Param surname query string false "Фамилия"
// @Param patronymic query string false "Отчество"
// @Param gender query string false "Пол"
// @Param nation query string false "Национальность"
// @Param min_age query int false "Минимальный возраст"
// @Param max_age query int false "Максимальный возраст"
// @Param limit query int false "Количество записей для возврата"
// @Param offset query int false "Смещение для пагинации"
// @Success 200 {array} entity.Person
// @Failure 400 {object} map[string]string "Неверные параметры фильтра"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /persons [get]
func GetPeople(c *gin.Context) {
	var filter entity.PersonFilter
	if err := c.ShouldBindQuery(&filter); err != nil || !filter.Valid() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверные параметры фильтра"})
		return
	}

	peoples, err := personsdb.Get(&filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, peoples)
	}
}

// @Summary Добавить нового человека
// @Description Создаёт новую запись о человеке
// @Tags persons
// @Accept json
// @Produce json
// @Param person body entity.PersonRequest true "Данные нового человека"
// @Success 201 {object} entity.Person
// @Failure 400 {object} map[string]string "Неверные данные"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /persons [post]
func AddPerson(c *gin.Context) {
	var newPersonReq entity.PersonRequest
	if err := c.ShouldBindJSON(&newPersonReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newPerson := newPersonReq.ToPerson()
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

// @Summary Обновить информацию о человеке
// @Description Обновляет данные существующего человека по ID
// @Tags persons
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Param person body entity.PersonRequest true "Обновлённые данные"
// @Success 200 {object} entity.Person
// @Failure 400 {object} map[string]string "Неверный ID или тело запроса"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /persons/{id} [put]
func UpdatePerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	var updatedReq entity.PersonRequest
	if err := c.ShouldBindJSON(&updatedReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated := updatedReq.ToPerson()
	refinePerson(&updated)

	updated.Id = id
	err = personsdb.Update(&updated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, updated)
	}
}

// @Summary Удалить человека
// @Description Удаляет человека по ID
// @Tags persons
// @Accept json
// @Produce json
// @Param id path int true "ID человека"
// @Success 204 "Успешное удаление"
// @Failure 400 {object} map[string]string "Неверный ID"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /persons/{id} [delete]
func DeletePerson(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный id"})
		return
	}

	err = personsdb.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.Status(http.StatusNoContent)
	}
}
