package config

import (
	"github.com/gin-gonic/gin"
)

func SetGinRouter() *gin.Engine {
	r := gin.Default()
	return r
}
