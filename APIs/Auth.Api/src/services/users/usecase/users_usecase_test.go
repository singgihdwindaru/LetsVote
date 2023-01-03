package usecase

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	mock_models "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/mock"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
	"github.com/stretchr/testify/suite"
)

type userSuiteTest struct {
	suite.Suite
	accountRepoMock *mock_models.MockIUserMysqlRepository
}

func (s *userSuiteTest) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	defer ctrl.Finish()

	s.accountRepoMock = mock_models.NewMockIUserMysqlRepository(ctrl)
	NewUsersUsecase(s.accountRepoMock)
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(userSuiteTest))
}

func (s *userSuiteTest) TestCreateUser() {
	// TODO
}

func TestCreateBlock(t *testing.T) {
	nikStr := "123"
	blockchain := models.NewBlockchain(nikStr, []byte{})
	hash := fmt.Sprintf("%x", blockchain.Blocks[0].CurrentBlockHash)

	t.Logf("\nhash : %v", hash)
}
