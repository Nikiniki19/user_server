package handler

import (
	"context"
	"errors"
	"log"
	"userservice/database"
	"userservice/models"
	"userservice/proto"
)

type CreateUser struct {
	proto.Client1RequestServer
}

func (c *CreateUser) CreateUser(ctx context.Context, userData *proto.UserDetails) (*proto.UserResponse1, error) {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}
	userDetails := models.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: userData.Password,
	}
	result := db.Create(&userDetails)
	if result.Error != nil && result.RowsAffected == 0 {
		return &proto.UserResponse1{}, errors.New("could not create the user")
	}
	return &proto.UserResponse1{
		Username: userDetails.Username,
		Password: userDetails.Password,
		Email: userDetails.Email,
	}, nil
}
