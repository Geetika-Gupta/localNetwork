package main

import (
	cv "github.com/ajira/project/custom_validator"
	"github.com/ajira/project/network"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	cv.RegisterValidations()
	network.InitializeRoutes(router)
	router.Run()
}
