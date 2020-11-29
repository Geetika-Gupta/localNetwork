package util

//Response handler.
func BuildResponse(data interface{}, err error, context *gin.Context) {
	var status int
	if err == nil {
		status = http.StatusOK
	} else {
		status = http.StatusBadRequest
	}
	context.JSON(
		status,
		gin.H{
			"status": status,
			"result": data,
			"error":  err,
		},
	)
}

func getPath(source, destination string) (path []string) {
	if source == destination {
		path = append(path, source)
		return
	}
	if _, ex := constant.COMPUTER[source]; ex {
		for _, val := range constant.CONNECTION[source] {
			getPath(val, destination)
		}
	}
	return
}