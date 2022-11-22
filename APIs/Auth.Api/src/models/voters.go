package models

import "context"

type (
	IVoterMysqlRepository interface {
		CreateVoter(ctx context.Context, nik, nonce int64, previousHash, hash string) error
		GetVotersByNIK(ctx context.Context, nik int64) (*Voter, error)
	}
	IVoterUsecase interface {
		CreateVoter(ctx context.Context, request CreateVoterRequest) (*CreateVoterResponse, error)
		GetVotersByNIK(ctx context.Context, nik int64) (*Voter, error)
	}
)
type (
	Voter struct {
		UserId       int64  `json:"userId"`
		Nonce        int64  `json:"nonce"`
		PreviousHash string `json:"previousHash"`
		Hash         string `json:"hash"`
	}

	CreateVoterRequest struct {
		NIK int64 `binding:"required" json:"nik"`
	}
	CreateVoterResponse struct {
		RegistrationNumber string `binding:"required" json:"registrationNumber"`
	}
)
