package models

import "gorm.io/gorm"

type Employee struct {
	FullName    string `json:"fullName"`    //ФИО
	Email       string `json:"email"`       //Е-mail
	NumberPhone string `json:"numberPhone"` //Телефон
	Job         string `json:"job"`         //Место работ
	Position    string `json:"position"`    //Должность
	Ranks       string `json:"ranks"`       //Звания
	Report      string `json:"report"`      //Название доклада
	FileName    string `json:"fileName"`
	TypeFile    string `json:"typeFile"`
	gorm.Model
}
type Sign struct {
	ID       int64  `json:"id" gorm:"primaryKey"` // uuid
	Name     string `json:"name"`
	Login    string `json:"login" gorm:"unique"`
	Password string `json:"password,omitempty"`
}

type Files struct {
	ID       int64  `json:"ID" gorm:"primaryKay"`
	FileName string `json:"fileName"`
	TypeFile string `json:"typeFile"`
}

type Docx struct {
	FileName  string `json:"fileName"`
	Positions string `json:"positions"`
}
