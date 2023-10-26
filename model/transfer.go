package model

type Transfer struct {
	Id          string `json:"id" binding:"required"`
	SenderID    string `json:"sender_id" binding:"required"`
	ReceiverID  string `json:"receiver_id" binding:"required"`
	Amount      string `json:"amount"`
	Description string `json:"description" binding:"max=255"`
}
