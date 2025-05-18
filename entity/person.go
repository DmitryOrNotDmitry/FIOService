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

type PersonRequest struct {
	Name       string `json:"name" example:"Dmitriy"`
	Surname    string `json:"surname" example:"Ushakov"`
	Patronymic string `json:"patronymic,omitempty" example:"Vasilevich"`
}

func (pr *PersonRequest) ToPerson() Person {
	return Person{Name: pr.Name, Surname: pr.Surname, Patronymic: pr.Patronymic}
}

type PersonFilter struct {
	Name       *string `form:"name"`
	Surname    *string `form:"surname"`
	Patronymic *string `form:"patronymic"`
	Gender     *string `form:"gender"`
	Nation     *string `form:"nation"`
	MinAge     *int    `form:"min_age"`
	MaxAge     *int    `form:"max_age"`

	Limit  int `form:"limit"`
	Offset int `form:"offset"`
}

func (pf *PersonFilter) Valid() bool {
	if pf.MinAge != nil && pf.MaxAge != nil {
		return *pf.MinAge <= *pf.MaxAge
	}

	return true
}
