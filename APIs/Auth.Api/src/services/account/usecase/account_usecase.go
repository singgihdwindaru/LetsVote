package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type accountUsecase struct {
	accountMysqlRepository models.IUserMysqlRepository
	voterMysqlRepository   models.IVoterMysqlRepository
}

func NewAccountUsecase(accountMysqlRepository models.IUserMysqlRepository, voterMysqlRepository models.IVoterMysqlRepository) models.IUserUsecase {
	return &accountUsecase{
		accountMysqlRepository: accountMysqlRepository,
		voterMysqlRepository:   voterMysqlRepository,
	}
}

// CreateParticipant implements models.IUserUsecase
func (u *accountUsecase) CreateVoter(ctx context.Context, request models.CreateVoterRequest) (*models.CreateVoterResponse, error) {
	user, err := u.accountMysqlRepository.GetUserByNIK(ctx, request.NIK)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	voter, err := u.voterMysqlRepository.GetVotersByNIK(ctx, request.NIK)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if voter != nil {
		err = errors.New("nik already registered")
		log.Println(err.Error())
		return nil, err
	}
	nikStr := strconv.FormatInt(int64(request.NIK), 10)
	blockchain := models.NewBlockchain(nikStr, []byte{})
	userBlock := fmt.Sprintf("%x", blockchain.Blocks[0].MyBlockHash)
	if userBlock != user.Hash {
		err = errors.New("user hash had been compromised")
		log.Println(err.Error())
		return nil, err
	}
	nonce := 1345 // TODO Generate this nonce (number only once)
	nonceStr := strconv.FormatInt(int64(nonce), 10)
	blockchain.AddBlock(nonceStr)
	participantBlock := fmt.Sprintf("%x", blockchain.Blocks[1].MyBlockHash)

	err = u.voterMysqlRepository.CreateVoter(ctx, request.NIK, int64(nonce), userBlock, participantBlock)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := models.CreateVoterResponse{
		RegistrationNumber: nonceStr,
	}
	return &result, err

}

// CreateUser implements models.IUserUsecase
func (u *accountUsecase) CreateUser(ctx context.Context, request models.CreateUserRequest) (*models.CreateUserResponse, error) {
	user, err := u.accountMysqlRepository.GetUserByNIK(ctx, request.NIK)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if user != nil {
		err = errors.New("nik already exists")
		log.Println(err.Error())
		return nil, err
	}
	guid := uuid.New().String()
	metadata, _ := json.Marshal(request)
	nikStr := strconv.FormatInt(int64(request.NIK), 10)
	blockchain := models.NewBlockchain(nikStr, []byte{})
	hash := fmt.Sprintf("%x", blockchain.Blocks[0].MyBlockHash)

	err = u.accountMysqlRepository.InsertUser(ctx, guid, string(metadata), hash)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := &models.CreateUserResponse{
		Data: request,
	}

	return result, nil
}

// SignIn implements models.IUserUsecase
func (u *accountUsecase) SignIn(ctx context.Context, request models.SignInRequest) (*models.SignInResponse, error) {
	panic("unimplemented")
}
