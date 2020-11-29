package network

//InitializeRoutes defines endpoints.
func InitializeRoutes(router *gin.Engine) {
	network := router.Group("/ajira/process")
	{
		network.POST("/devices", addRequest())
		network.POST("/connections", connectDevices())
		network.GET("/devices", listDevices())
		network.GET("/info-routes", routeInfo())
		network.PUT("/devices/:name/strength", updateStrength())
	}
}

//will add new devices 
func addRequest() func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			req model.AddRequest
			err error
		)
		if req, err = util.VAddRequest(context); err == nil {
			data, err = util.PAddRequest(req)
		}
		util.BuildResponse(data, err, context)
	}
}

//will connect devices
func connectDevices() func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			req model.ConnectDevice
			err error
		)
		if req, err = util.VConnectDevices(context); err == nil {
			data, err = util.PConnectDevices(req)
		}
		util.BuildResponse(data, err, context)
	}
}

//will list devices
func listDevices() func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data model.ListDevices
			err error
		)
		data, err = util.PListDevices()
		util.BuildResponse(data, err, context)
	}
}

//will connect devices
func updateStrength(name string) func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			req model.UpdateStrength
			err error
		)
		if req, err = util.VUpdateStrength(context); err == nil {
			data, err = util.PUpdateStrength(name, req)
		}
		util.BuildResponse(data, err, context)
	}
}

//will give route info
func routeInfo(name string) func(context *gin.Context) {
	return func(context *gin.Context) {
		var (
			data string
			err error
		)
		data, err = util.PRouteInfo(context.Param("from"), context.Param("to"))
		util.BuildResponse(data, err, context)
	}
}