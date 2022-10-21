package repository

import (
	"errors"
	"final_assignment/app/models"
	"final_assignment/app/resource"
	"final_assignment/config"
)

type CommentRepository interface {
	NewComment(Comment *models.Comment, createData resource.NewComment) error
	GetComments(Comments *[]models.Comment, userId uint) error
	EditComment(Comment *models.Comment, createData resource.EditComment, userId uint) error
	RemoveComment(userId uint, id uint) error
}

func NewCommentRepository() CommentRepository {
	return &dbConnection{
		connection: config.Connect(),
	}
}

func (db *dbConnection) NewComment(Comment *models.Comment, createData resource.NewComment) error {
	Comment.Message = createData.Message
	Comment.PhotoID = createData.PhotoID
	err := db.connection.Save(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) EditComment(Comment *models.Comment, createData resource.EditComment, userId uint) error {
	db.connection.Model(&Comment).Where("id = ?", Comment.ID).First(&Comment)
	if Comment.ID != 0 {
		if Comment.UserID != userId {
			return errors.New("You not authorized!")
		}
	}
	Comment.Message = createData.Message
	err := db.connection.Save(Comment).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetComments(Comments *[]models.Comment, userId uint) error {
	db.connection.Model(Comments).Where("user_id = ?", userId).Preload("User").Preload("Photo").Find(Comments)
	return nil
}

func (db *dbConnection) RemoveComment(userId uint, id uint) error {
	var Comment models.Comment
	db.connection.Model(Comment).Where("id = ?", id).Find(&Comment)
	if Comment.ID != 0 {
		if Comment.UserID != userId {
			return errors.New("You not authorized!")
		}
	} else {
		return errors.New("Data not found")
	}
	err := db.connection.Unscoped().Delete(&Comment).Error
	if err != nil {
		return err
	}
	return nil
}
