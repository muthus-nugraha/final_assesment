package repository

import (
	"errors"
	"final_assignment/app/models"
	"final_assignment/app/resource"
	"final_assignment/app/utils"
	"final_assignment/config"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Signup(User *models.User, createData resource.NewUser) error
	Signin(email string, password string) (string, error)
	EditUser(User *models.User, createData resource.EditUser) error
	RemoveUser(id int) error
	HashPassword(text string) string
}

func NewUserRepository() UserRepository {
	return &dbConnection{
		connection: config.Connect(),
	}
}

func (db *dbConnection) Signup(User *models.User, createData resource.NewUser) error {
	User.Username = createData.Username
	User.Email = createData.Email
	User.Age = createData.Age
	User.Password = db.HashPassword(createData.Password)
	err := db.connection.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) Signin(email string, password string) (string, error) {
	var User models.User
	db.connection.Model(&User).Where("email = ?", email).Find(&User)
	if User.ID != 0 {
		if db.CheckPassword(password, User.Password) {
			token, err := utils.GenerateToken(User.ID)
			if err != nil {
				return "", errors.New("Please try again, for generate password")
			}
			return token, nil
		} else {
			return "", errors.New("Email or password not match")
		}
	}
	return "", errors.New("Email or password not match")
}

func (db *dbConnection) EditUser(User *models.User, createData resource.EditUser) error {
	db.connection.Model(&User).Where("id = ?", User.ID).Find(&User)

	User.Username = createData.Username
	User.Email = createData.Email
	err := db.connection.Save(User).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) GetUsers() ([]models.User, error, int64) {
	var User []models.User
	var count int64
	connection := db.connection.Model(&User).Preload("Items").Find(&User)
	err := connection.Error
	if err != nil {
		return User, err, 0
	}
	db.connection.Model(User).Count(&count)
	return User, nil, count
}

func (db *dbConnection) RemoveUser(id int) error {
	var User models.User
	err := db.connection.Unscoped().Delete(&User, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConnection) HashPassword(text string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(text), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (db *dbConnection) CheckPassword(password string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == nil {
		return true
	}
	return false
}
