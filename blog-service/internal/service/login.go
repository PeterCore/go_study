package service

type LoginRequset struct {
	UserName string `json:"username" binding:"max=100"`
	Password string `json:"pasword" binding:"max=16"`
}
