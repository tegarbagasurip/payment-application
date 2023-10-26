package model

type Profile struct {
	Id      string `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required,max=100"`
	Address string `json:"address" binding:"required,max=100"`
	Phone   string `json:"phone" binding:"required,max=15"`
	Balance string `json:"balance" binding:"required,max=1000000"`
}