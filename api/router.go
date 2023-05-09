package api

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes(router *gin.Engine) {

	router.GET("/", hello)

	router.GET("/lets-go", boardInit)

	router.POST("/play", makeAMove)

	/*
		Tests
	*/
	router.POST("/diagonal-test", diagonalTest)

}
