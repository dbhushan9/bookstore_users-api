package services

import (
	"github.com/dbhushan9/bookstore_users-api/domain/users"
	"github.com/dbhushan9/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	err := user.Save()
	return &user, err
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := users.User{Id: userId}
	err := result.Get()
	return &result, err
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {

	current, err := GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}

func DeleteUser(user_id int64) *errors.RestErr {
	user := users.User{Id: user_id}
	return user.Delete()
}
