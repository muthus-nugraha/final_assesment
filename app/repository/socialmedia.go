package repository

import (
	"errors"
	"final_assignment/app/models"
	"final_assignment/app/resource"
	"final_assignment/config"
	
)

type SocialMediaRepository interface {
	NewSocialMedia(SocialMedia *models.SocialMedia, createData resource.NewSocialMedia) error
	GetSocialMedia(SocialMedias *[]models.SocialMedia, userId uint) error
	EditSocialMedia(SocialMedia *models.SocialMedia, createData resource.EditSocialMedia, userId uint) error
	RemoveSocialMedia(userId uint, id uint) error
}

func NewSocialMediaRepository() SocialMediaRepository {
	return &dbConnection{
		connection: config.Connect(),
	}
}

func (db *dbConnection) NewSocialMedia(SocialMedia *models.SocialMedia, createData resource.NewSocialMedia) error {
	SocialMedia.Name = createData.Name
	SocialMedia.SocialMedialUrl = createData.SocialMedialUrl
	err := db.connection.Save(SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) EditSocialMedia(SocialMedia *models.SocialMedia, createData resource.EditSocialMedia, userId uint) error {
	db.connection.Model(&SocialMedia).Where("id = ?", SocialMedia.ID).First(&SocialMedia)
	if SocialMedia.ID != 0 {
		if SocialMedia.UserID != userId {
			return errors.New("You not authorized!")
		}
	}
	SocialMedia.Name = createData.Name
	SocialMedia.SocialMedialUrl = createData.SocialMedialUrl
	err := db.connection.Save(SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetSocialMedia(SocialMedias *[]models.SocialMedia, userId uint) error {
	db.connection.Model(SocialMedias).Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(SocialMedias)
	return nil
}

func (db *dbConnection) RemoveSocialMedia(userId uint, id uint) error {
	var SocialMedia models.SocialMedia
	db.connection.Model(SocialMedia).Where("id = ?", id).Find(&SocialMedia)
	if SocialMedia.ID != 0 {
		if SocialMedia.UserID != userId {
			return errors.New("You not authorized!")
		}
	} else {
		return errors.New("Data not found")
	}
	err := db.connection.Unscoped().Delete(&SocialMedia).Error
	if err != nil {
		return err
	}
	return nil
}
