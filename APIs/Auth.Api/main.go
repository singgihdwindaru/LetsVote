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
	httpUsers "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/users/controller/http"
	usersRepo "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/users/repository/mysql"
	usersUc "github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/services/users/usecase"
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
	usersMysqlRepo = usersRepo.NewUsersMysqlRepository(DB)
	voterMysqlRepo = voterRepo.NewVoterMysqlRepository(DB)
}

func initUsecase() {
	userUsecase = usersUc.NewUsersUsecase(usersMysqlRepo)
	voterUsecase = voterUc.NewVoterUsecase(usersMysqlRepo, voterMysqlRepo)
}
func SetupApi() *gin.Engine {
	initRepo()
	initUsecase()
	r := gin.Default()
	httpUsers.NewUsersController(r, userUsecase)
	httpVoter.NewVoterController(r, voterUsecase)
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
