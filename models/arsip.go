package models

import "gorm.io/gorm"

type Arsip struct {
	gorm.Model
	Judul string `json:"judul"`
	Isi   string `json:"isi"`
}
