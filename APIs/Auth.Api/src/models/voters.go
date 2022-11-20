package models

import "context"

type (
	IVoterMysqlRepository interface {
		CreateVoter(ctx context.Context, nik, nonce int64, previousHash, hash string) error
		GetVotersByNIK(ctx context.Context, nik int64) (*User, error)
	}
)
type (
	CreateVoterRequest struct {
		NIK int64 `binding:"required" json:"nik"`
	}
	CreateVoterResponse struct {
		RegistrationNumber string `binding:"required" json:"registrationNumber"`
	}
)
