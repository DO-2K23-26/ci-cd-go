package models

type City struct {
	ID             uint    `json:"id" gorm:"primaryKey;not null"`
	DepartmentCode string  `json:"department_code"`
	InseeCode      string  `json:"insee_code"`
	ZipCode        string  `json:"zip_code"`
	Name           string  `json:"name"`
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
}
