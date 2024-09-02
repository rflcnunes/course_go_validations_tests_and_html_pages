package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name string `json:"name" gorm:"column:name"`
	CPF  string `json:"cpf" gorm:"column:cpf"`
	RG   string `json:"rg" gorm:"column:rg"`
}
