package main

import (
	//"net/http"

	"github.com/gin-gonic/gin"

	"go-projects/connect-4/api"
)

func main() {

	router := gin.Default()

	api.InitRoutes(router)

	router.Run("localhost:1000")
}
