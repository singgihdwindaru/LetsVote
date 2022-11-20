package controller_http

import (
	"github.com/gin-gonic/gin"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type accountController struct {
	authUsecase models.IUserUsecase
}

func NewAuthController(r *gin.Engine, authUsecase models.IUserUsecase) {
	ctrl := &accountController{
		authUsecase: authUsecase,
	}

	r.POST("/account/user", ctrl.CreateUser)
	r.POST("/account/signin", ctrl.SignIn)
	r.POST("/account/participant", ctrl.CreateParticipant)
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

}
func (c *accountController) CreateParticipant(g *gin.Context) {

}
