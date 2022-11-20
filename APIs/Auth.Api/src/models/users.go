package models

import (
	"context"
)

type (
	IUserRepository interface {
		InsertUser(ctx context.Context, email, name, password string) error
		GetUserByNIK(ctx context.Context, nik string) (*User, error)
	}
	IUserUsecase interface {
		CreateUser(ctx context.Context, request CreateUserRequest) (*CreateUserResponse, error)
		SignIn(ctx context.Context, request SignInRequest) (*SignInResponse, error)
		CreateParticipant(ctx context.Context, request CreateParticipantRequest) (*CreateParticipantResponse, error)
	}
)
type (
	User struct {
		Id       int64  `json:"id"`
		Guid     string `json:"guid"`
		Metadata string `json:"metadata"`
		Hash     string `json:"hash"`
	}

	CreateUserRequest struct {
		NIK       int64  `json:"nik"`
		FullName  string `json:"fullName"`
		Address   string `json:"address"`
		Gender    string `json:"gender"`
		BloodType string `json:"bloodType"`
	}

	CreateUserResponse struct {
		Data CreateUserRequest `json:"data"`
	}

	SignInRequest struct {
		RegistrationNumber string `binding:"required" json:"registrationNumber"`
	}

	SignInResponse struct {
		Token string `binding:"required" json:"token"`
		Hash  string `binding:"required" json:"hash"`
	}

	CreateParticipantRequest struct {
		NIK string `binding:"required" json:"registrationNumber"`
	}
	CreateParticipantResponse struct {
		RegistrationNumber string `binding:"required" json:"registrationNumber"`
	}
)
