package model

//AddRequest : add request model
type AddRequest struct {
	Type string `json:"type" binding:"required,devicetype"`
	Name string `json:"name" binding:"required"`
}

//ConnectDevice : connect device request model
type ConnectDevice struct {
	Source string `json:"source" binding:"required"`
	Target []string `json:"target"`
}

//UpdateStrength : udate strength request model
type UpdateStrength struct {
	Strength int `json:"value"`
}