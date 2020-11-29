package util

//PAddRequest : will process add device request
func PAddRequest(req model.AddRequest) (data string, err error) {
	if strings.ToLower(req.Type) == "computer" {
		if _, exists := constant.COMPUTER[req.Name]; !exists {
			constant.COMPUTER[req.Name] = 5 //default strength is 5
		} else {
			err = perror.CustomError("DEVICE ALREADY EXISTS")
			return
		}
	} else {
		if _, exists := constant.REPEATER[req.Name]; !exists {
			constant.REPEATER[req.Name] = 0 //repeaters font have any strength
		} else {
			err = perror.CustomError("DEVICE ALREADY EXISTS")
			return
		}
	}
	data = "DEVICE ADDED SUCCESSFULLY!!!"
	return
}

//PConnectDevices : will process connect devices request
func PConnectDevices(req model.AddRequest) (data string, err error) {
	var exists bool
	if _, exists = constant.COMPUTER[req.Source]; !exists {
		_, exists = constant.REPEATER[req.Source]
	}
	if exists {
		//check if any current connection for device exists
		if currConn, ex := constant.CONNECTION[req.Source]; !ex {
			currConn = make(map[int]struct{})
		}
		for _, val := range req.Target {
			if val == req.Source {
				err = perror.CustomError("DEVICE CANNOT CONNECT TO ITSELF")
				return
			}
			var hasDevice bool
			if _, ex := constant.COMPUTER[val]; ex {
				hasDevice = true
			} else if _, ex = constant.REPEATER[val]; ex {
				hasDevice = true
			}
			if hasDevice {
				if _, exists = currConn[val]; !exists {
					currConn[val] = struct{}{}
				}
			} else {
				err = perror.CustomError("DEVICE DOESN'T EXISTS")
				return
			}
		}
		constant.CONNECTION[req.Source] = currConn
	} else {
		err = perror.CustomError("NO SOURCE FOUND ERROR")
	}
	data = "DEVICES CONNECTION MAKE SUCCESSFULLY!!!"
	return
}

//PListDevices : will process list devices
func PListDevices() (data model.ListDevices, err error) {
	data = make([]model.ListDevices, 0)
	for _, val := range constant.COMPUTER {
		data = append(data, model.Device{
			Type: "COMPUTER",
			Name: val,
		})
	}
	for _, val := range constant.REPEATER {
		data = append(data, model.Device{
			Type: "REPEATER",
			Name: val,
		})
	}
	return
}

//PUpdateStrength : will process update device strength request
func PUpdateStrength(name string, req model.AddRequest) (data string, err error) {
	if _, exists := constant.COMPUTER[name]; exists {
		constant.COMPUTER[name] = req.Strength
		if req.Strength == 0 {
			constant.COMPUTER[name] = 5
		} 
		data = "STRENGTH UPDATED FOR DEVICE SUCCESSFULLY!!!"
	} else {
		err = perr.CustomError("NO COMPUTER DEVICE FOUND TO UPDATE STRENGTH")
	}
	return
}

//PRouteInfo : will process route-info get request
func PRouteInfo(source, destination string) (data string, err error) {
	if _, exists := constant.COMPUTER[source]; !exists {
		err = perror.CustomError("SOURCE SHOULD BE COMPUTER")
		return
	}
	if _, exists := constant.COMPUTER[destination]; !exists {
		err = perror.CustomError("DESTINATION HAVE TO BE COMPUTER")
		return
	}
	getPath(source, destination)
	return
}