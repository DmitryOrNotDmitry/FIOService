package entity

type Person struct {
	Id         int64  `json:"id,omitempty" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic,omitempty"`

	Age    int    `json:"age,omitempty"`
	Gender string `json:"gender,omitempty"`
	Nation string `json:"nation,omitempty"`
}

func (Person) TableName() string {
	return "persons"
}
