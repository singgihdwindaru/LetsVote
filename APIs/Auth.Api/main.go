package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/config"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
	controller "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/controller/http"
	accountRepo "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/repository/mysql"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/account/usecase"
	voterRepo "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/voter/repository/mysql"
)

var (
	DB                   *sql.DB
	usersMysqlRepo       models.IUserMysqlRepository
	voterMysqlRepository models.IVoterMysqlRepository

	userUsecase models.IUserUsecase
)

func initRepo() {
	usersMysqlRepo = accountRepo.NewAccountMysqlRepository(DB)
	voterMysqlRepository = voterRepo.NewVoterMysqlRepository(DB)
}

func initUsecase() {
	userUsecase = usecase.NewAccountUsecase(usersMysqlRepo, voterMysqlRepository)
}
func SetupApi() *gin.Engine {
	initRepo()
	initUsecase()
	r := gin.Default()
	controller.NewAccountController(r, userUsecase)

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
