package handler

import (
	"context"
	"errors"
	"log"
	"strconv"
	"userservice/database"
	"userservice/models"
	"userservice/proto"

	"gorm.io/gorm"
)

type FetchUser struct {
	proto.Client2RequestServer
}

func (f *FetchUser) FetchUser(ctx context.Context, userID *proto.Id) (*proto.UserResponse2, error) {
	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Fatalf("could not connect to db %v", err)
	}
	u64, err := strconv.ParseUint(userID.Id, 10, 32)
	if err != nil {
		return nil, errors.New("could not convert string to uint")
	}
	useradata := models.User{
		Model: gorm.Model{
			ID: uint(u64),
		},
	}
	result := db.Where("id = ?", useradata.ID).First(&useradata)
	if result.Error != nil && result.RowsAffected == 0 {
		return &proto.UserResponse2{}, errors.New("could not fetch the user")
	}
	return &proto.UserResponse2{
		Username: useradata.Username,
		Email:    useradata.Email,
	}, nil
}
