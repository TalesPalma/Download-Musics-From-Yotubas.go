package models

type Link struct {
	Link string `json:"link" binding:"required"`
}
