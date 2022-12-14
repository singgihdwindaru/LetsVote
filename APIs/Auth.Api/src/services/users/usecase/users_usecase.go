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

type usersUsecase struct {
	accountMysqlRepository models.IUserMysqlRepository
}

func NewUsersUsecase(accountMysqlRepository models.IUserMysqlRepository) models.IUserUsecase {
	return &usersUsecase{
		accountMysqlRepository: accountMysqlRepository,
	}
}

// CreateUser implements models.IUserUsecase
func (u *usersUsecase) CreateUser(ctx context.Context, request models.CreateUserRequest) (*models.CreateUserResponse, error) {
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
	hash := fmt.Sprintf("%x", blockchain.Blocks[0].CurrentBlockHash)

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
func (u *usersUsecase) SignIn(ctx context.Context, request models.SignInRequest) (*models.SignInResponse, error) {
	panic("unimplemented")
}
