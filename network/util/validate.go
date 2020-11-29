package util

import (
	"github.com/gin-gonic/gin"
)

//VAddRequest : will validate add device request
func VAddRequest(context *gin.Context) (req model.AddRequest, err error) {
	err := context.ShouldBindJSON(&req)
	return
}

//VConnectDevices : will validate connect device request
func VConnectDevices(context *gin.Context) (req model.ConnectDevice, err error) {
	err := context.ShouldBindJSON(&req)
	return
}

//VUpdateStrength : will validate strength device request
func VUpdateStrength(context *gin.Context) (req model.UpdateStrength, err error) {
	err := context.ShouldBindJSON(&req)
	return
}