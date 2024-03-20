package models

type Terminal struct {
	Id   string `json:"id" gorm:"primaryKey;type:CHAR(36)"`
	Name string `json:"name" gorm:"size:100;unique"`
}
