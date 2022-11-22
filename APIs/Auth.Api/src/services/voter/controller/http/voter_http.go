package controller_http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/singgihdwindaru/LetsVote/APIs/Auth.Api/src/models"
)

type voterController struct {
	voterUsecase models.IVoterUsecase
}

func NewAccountController(r *gin.Engine, voterUsecase models.IVoterUsecase) {
	ctrl := &voterController{
		voterUsecase: voterUsecase,
	}

	r.GET("/voter", ctrl.GetVoterByNik)
	r.POST("/voter", ctrl.CreateVoter)
}

func (c *voterController) GetVoterByNik(g *gin.Context) {
	nik, err := strconv.ParseInt(g.Query("nik"), 10, 64)
	if err != nil {
		log.Println(err.Error())
		// TODO create utils for mapping statuscode and message
		g.JSON(int(http.StatusUnprocessableEntity), models.HttpResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}
	result, err := c.voterUsecase.GetVotersByNIK(g.Request.Context(), int64(nik))
	if err != nil {
		log.Println(err.Error())
		// TODO create utils for mapping statuscode and message
		g.JSON(http.StatusInternalServerError, models.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, models.HttpResponse(http.StatusOK, "Success", result))
}

func (c *voterController) CreateVoter(g *gin.Context) {
	request := models.CreateVoterRequest{}

	if err := g.BindJSON(&request); err != nil {
		log.Println(err.Error())
		g.JSON(http.StatusBadRequest, models.HttpResponse(http.StatusBadRequest, "Invalid request", nil))
		return
	}

	result, err := c.voterUsecase.CreateVoter(g.Request.Context(), request)
	if err != nil {
		log.Println(err.Error())
		// TODO create utils for mapping statuscode and message
		g.JSON(http.StatusInternalServerError, models.HttpResponse(http.StatusInternalServerError, err.Error(), nil))
		return
	}
	g.JSON(http.StatusOK, models.HttpResponse(http.StatusOK, "Success", result))
}
