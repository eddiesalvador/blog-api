package models

import (
	"gorm.io/gorm"
	_"time"
)

type User struct {
	gorm.Model
	/*
	ID
	CreatedAt
	UpdatedAt
	DeletedAt
	*/
	FirstName		string
	MiddleName  	string
	LastName		string
	Mobile			string
	Email 			string
	PasswordHash	string
	Intro			string
	Profile			string
}

//create a user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get all users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetUser(db *gorm.DB, User *User, id int) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)
	return nil
}

//delete user
func DeleteUser(db *gorm.DB, User *User, id int) (err error) {
	db.Where("id = ?", id).Delete(User)
	return nil
}