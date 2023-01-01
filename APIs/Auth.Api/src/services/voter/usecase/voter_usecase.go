package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/common"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type voterUsecase struct {
	userMysqlRepository  models.IUserMysqlRepository
	voterMysqlRepository models.IVoterMysqlRepository
}

func NewVoterUsecase(userMysqlRepository models.IUserMysqlRepository, voterMysqlRepository models.IVoterMysqlRepository) models.IVoterUsecase {
	return &voterUsecase{
		userMysqlRepository:  userMysqlRepository,
		voterMysqlRepository: voterMysqlRepository,
	}
}

// CreateVoter implements models.IVoterUsecase
func (u *voterUsecase) CreateVoter(ctx context.Context, request models.CreateVoterRequest) (*models.CreateVoterResponse, error) {
	user, err := u.userMysqlRepository.GetUserByNIK(ctx, request.NIK)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user == nil {
		err = errors.New("nik not registered yet")
		log.Println(err.Error())
		return nil, err
	}
	voter, err := u.voterMysqlRepository.GetVotersByNIK(ctx, request.NIK)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if voter != nil {
		err = errors.New("nik became voter already")
		log.Println(err.Error())
		return nil, err
	}
	nikStr := strconv.FormatInt(int64(request.NIK), 10)
	blockchain := models.NewBlockchain(nikStr, []byte{})
	userBlock := fmt.Sprintf("%x", blockchain.Blocks[0].CurrentBlockHash)
	if userBlock != user.Hash {
		err = errors.New("registered user's hash had been compromised")
		log.Println(err.Error())
		return nil, err
	}
	nonce, err := common.RandomIntGenerator(4)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	nonceStr := strconv.FormatInt(int64(nonce), 10)
	blockchain.AddBlock(nonceStr)
	voterBlock := fmt.Sprintf("%x", blockchain.Blocks[1].CurrentBlockHash)

	err = u.voterMysqlRepository.CreateVoter(ctx, request.NIK, int64(nonce), userBlock, voterBlock)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := models.CreateVoterResponse{
		RegistrationNumber: nonceStr,
	}
	return &result, err
}

// GetVotersByNIK implements models.IVoterUsecase
func (u *voterUsecase) GetVotersByNIK(ctx context.Context, nik int64) (*models.Voter, error) {
	voter, err := u.voterMysqlRepository.GetVotersByNIK(ctx, nik)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return voter, nil
}
