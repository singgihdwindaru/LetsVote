package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
	httpAccount "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/controller/http"
	accountRepo "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/repository/mysql"
	accountUc "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/usecase"
	httpVoter "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/voter/controller/http"
	voterRepo "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/voter/repository/mysql"
	voterUc "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/voter/usecase"
)

var (
	DB             *sql.DB
	usersMysqlRepo models.IUserMysqlRepository
	voterMysqlRepo models.IVoterMysqlRepository

	userUsecase  models.IUserUsecase
	voterUsecase models.IVoterUsecase
)

func init() {
	rand.Seed(time.Now().UnixNano())

}
func initRepo() {
	usersMysqlRepo = accountRepo.NewAccountMysqlRepository(DB)
	voterMysqlRepo = voterRepo.NewVoterMysqlRepository(DB)
}

func initUsecase() {
	userUsecase = accountUc.NewAccountUsecase(usersMysqlRepo)
	voterUsecase = voterUc.NewVoterUsecase(usersMysqlRepo, voterMysqlRepo)
}
func SetupApi() *gin.Engine {
	initRepo()
	initUsecase()
	r := gin.Default()
	httpAccount.NewAccountController(r, userUsecase)
	httpVoter.NewAccountController(r, voterUsecase)
	return r
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(err.Error())
	}

	DB = config.SetupDB()
	defer DB.Close()

	r := SetupApi()
	err = r.Run(":8080")
	log.Fatalf(err.Error())

}
