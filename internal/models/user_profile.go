package models

type UserProfile struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
	Group       string `json:"group"`
}
