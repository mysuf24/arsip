package services

import (
	"arsip/config"
	"arsip/models"
)

func GetAllArsip() ([]models.Arsip, error){
	var arsip []models.Arsip
	err:= config.DB.Find(&arsip).Error
	return arsip, err
}