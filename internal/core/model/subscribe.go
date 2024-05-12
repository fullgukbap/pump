package model

type Subscribe struct {
	ID int `json:"id"`

	UserID int `json:"userID"`
	ExrID  int `json:"exrID"`
}
