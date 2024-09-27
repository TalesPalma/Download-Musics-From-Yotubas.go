package models

type PlaylistLink struct {
	Link string `json:"playlist_url" binding:"required"`
}
