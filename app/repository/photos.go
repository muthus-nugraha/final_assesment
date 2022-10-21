package repository

import (
	"errors"
	"final_assignment/app/models"
	"final_assignment/app/resource"
	"final_assignment/config"
)

type PhotoRepository interface {
	NewPhoto(Photo *models.Photo, createData resource.NewPhoto) error
	GetPhotos(Photos *[]models.Photo, userId uint) error
	EditPhoto(Photo *models.Photo, createData resource.NewPhoto, userId uint) error
	RemovePhoto(userId uint, id uint) error
}

func NewPhotoRepository() PhotoRepository {
	return &dbConnection{
		connection: config.Connect(),
	}
}

func (db *dbConnection) NewPhoto(Photo *models.Photo, createData resource.NewPhoto) error {
	Photo.Title = createData.Title
	Photo.Caption = createData.Caption
	Photo.PhotoUrl = createData.PhotoUrl
	err := db.connection.Save(Photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) EditPhoto(Photo *models.Photo, createData resource.NewPhoto, userId uint) error {
	db.connection.Model(&Photo).Where("id = ?", Photo.ID).First(&Photo)
	if Photo.ID != 0 {
		if Photo.UserID != userId {
			return errors.New("You not authorized!")
		}
	}
	Photo.Title = createData.Title
	Photo.Caption = createData.Caption
	Photo.PhotoUrl = createData.PhotoUrl
	err := db.connection.Save(Photo).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetPhotos(Photos *[]models.Photo, userId uint) error {
	db.connection.Model(Photos).Where("user_id = ?", userId).Find(Photos)
	return nil
}

func (db *dbConnection) RemovePhoto(userId uint, id uint) error {
	var Photo models.Photo
	db.connection.Model(Photo).Where("id = ?", id).Find(&Photo)
	if Photo.ID != 0 {
		if Photo.UserID != userId {
			return errors.New("You not authorized!")
		}
	} else {
		return errors.New("Data not found")
	}
	err := db.connection.Unscoped().Delete(&Photo).Error
	if err != nil {
		return err
	}
	return nil
}
