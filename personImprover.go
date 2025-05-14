package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	ageApi    = "https://api.agify.io/?name=%s"
	genderApi = "https://api.genderize.io/?name=%s"
	nationApi = "https://api.nationalize.io/?name=%s"
)

func executeRequest(url string, data any) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Ошибка при запросе: %v", err)
	}
	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(data); err != nil {
		log.Fatalf("Ошибка при декодировании JSON: %v", err)
	}
}

type AgeData struct {
	Age int `json:"age"`
}

func GetAgeByName(name string) int {
	var decResp AgeData
	executeRequest(fmt.Sprintf(ageApi, name), &decResp)

	return decResp.Age
}

type GenderData struct {
	Gender string `json:"gender"`
}

func GetGenderByName(name string) string {
	var decResp GenderData
	executeRequest(fmt.Sprintf(genderApi, name), &decResp)

	return decResp.Gender
}

type Country struct {
	Id   string  `json:"country_id"`
	Prob float64 `json:"probability"`
}

type NationData struct {
	Countries []Country `json:"country"`
}

func GetNationByName(name string) string {
	var decResp NationData
	executeRequest(fmt.Sprintf(nationApi, name), &decResp)

	cId := ""
	maxProp := float64(0)
	for _, country := range decResp.Countries {
		if country.Prob > maxProp {
			maxProp = country.Prob
			cId = country.Id
		}
	}

	return cId
}
