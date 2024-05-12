package model

type User struct {
	ID            int    `json:"id"`
	UID           string `json:"uid"`
	Password      string `json:"password"`
	OauthProvider string `json:"oauthProvider"`
	OauthUID      int    `json:"oauthUID"`

	Notifications []Notification
	Subscribes    []Subscribe
}
