package main

import (
	"encoding/json"
	"fioservice/logger"
	"fmt"
	"net/http"
)

const (
	ageApi    = "https://api.agify.io/?name=%s"
	genderApi = "https://api.genderize.io/?name=%s"
	nationApi = "https://api.nationalize.io/?name=%s"
)

func executeRequest(url string, data any) error {
	logger.Log.Debugf("Выполнение GET запроса: %s", url)

	resp, err := http.Get(url)
	if err != nil {
		logger.Log.Infof("Ошибка при запросе к %s: %v", url, err)
		return err
	}
	defer resp.Body.Close()

	logger.Log.Debugf("Ответ от %s получен. Статус: %s", url, resp.Status)

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		logger.Log.Infof("Ошибка декодирования JSON от %s: %v", url, err)
		return err
	}

	logger.Log.Debugf("JSON успешно декодирован от %s", url)
	return nil
}

type AgeData struct {
	Age int `json:"age"`
}

func GetAgeByName(name string) (int, error) {
	logger.Log.Debugf("Получение возраста по имени: %s", name)

	var decResp AgeData
	err := executeRequest(fmt.Sprintf(ageApi, name), &decResp)
	if err != nil {
		return 0, err
	}

	logger.Log.Infof("Возраст для %s: %d", name, decResp.Age)
	return decResp.Age, nil
}

type GenderData struct {
	Gender string `json:"gender"`
}

func GetGenderByName(name string) (string, error) {
	logger.Log.Debugf("Получение пола по имени: %s", name)

	var decResp GenderData
	err := executeRequest(fmt.Sprintf(genderApi, name), &decResp)
	if err != nil {
		return "", err
	}

	logger.Log.Infof("Пол для %s: %s", name, decResp.Gender)
	return decResp.Gender, nil
}

type Country struct {
	Id   string  `json:"country_id"`
	Prob float64 `json:"probability"`
}

type NationData struct {
	Countries []Country `json:"country"`
}

func GetNationByName(name string) (string, error) {
	logger.Log.Debugf("Получение страны по имени: %s", name)

	var decResp NationData
	err := executeRequest(fmt.Sprintf(nationApi, name), &decResp)
	if err != nil {
		return "", err
	}

	cId := ""
	maxProp := float64(0)
	for _, country := range decResp.Countries {
		if country.Prob > maxProp {
			maxProp = country.Prob
			cId = country.Id
		}
	}

	logger.Log.Infof("Страна для %s: %s (%.2f)", name, cId, maxProp)
	return cId, nil
}
