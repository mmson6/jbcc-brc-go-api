package brcapiv1

type UserProfile struct {
	Group       string `json:"group"`       // ID of the assigned group
	DisplayName string `json:"displayName"` // Display name of the user
}
