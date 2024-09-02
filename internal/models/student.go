package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
	CPF  string `json:"cpf" gorm:"column:cpf;unique"`
	RG   string `json:"rg" gorm:"column:rg;unique"`
}
