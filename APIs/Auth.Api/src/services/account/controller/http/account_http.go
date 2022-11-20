package controller_http

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type accountController struct {
	accountUsecase models.IUserUsecase
}

func NewAccountController(r *gin.Engine, accountUsecase models.IUserUsecase) {
	ctrl := &accountController{
		accountUsecase: accountUsecase,
	}

	r.POST("/account/user", ctrl.CreateUser)
	r.POST("/account/signin", ctrl.SignIn)
	r.POST("/account/voter", ctrl.CreateVoter)
}

func (c *accountController) SignIn(g *gin.Context) {
	// request := models.LoginRequest{}

	// if err := g.BindJSON(&request); err != nil {
	// 	log.Println(err.Error())
	// 	g.JSON(http.StatusBadRequest, common.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
	// 	return
	// }

	// result, err := c.userUsecase.Login(g.Request.Context(), request)
	// if err != nil {
	// 	log.Println(err.Error())
	// 	g.JSON(http.StatusInternalServerError, common.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
	// 	return
	// }
	// g.JSON(http.StatusOK, common.HttpResponse(http.StatusOK, "Success SignIn", result))
}

func (c *accountController) CreateUser(g *gin.Context) {
	request := models.CreateUserRequest{}

	if err := g.BindJSON(&request); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, models.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	result, err := c.accountUsecase.CreateUser(g.Request.Context(), request)
	if err != nil {
		log.Println(err.Error())
		// TODO create utils for mapping statuscode and message
		g.JSON(http.StatusInternalServerError, models.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, models.HttpResponse(http.StatusOK, "Success", result.Data))

}
func (c *accountController) CreateVoter(g *gin.Context) {
	request := models.CreateVoterRequest{}

	if err := g.BindJSON(&request); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, models.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	result, err := c.accountUsecase.CreateVoter(g.Request.Context(), request)
	if err != nil {
		log.Println(err.Error())
		// TODO create utils for mapping statuscode and message
		g.JSON(http.StatusInternalServerError, models.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, models.HttpResponse(http.StatusOK, "Success", result))
}
