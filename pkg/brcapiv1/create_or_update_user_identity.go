package brcapiv1

type UserIdentity struct {
	ID          string `json:"id"`          // ID of the user
	GroupID     string `json:"groupId"`     // ID of the assigned group
	DisplayName string `json:"displayName"` // Display name of the user
}
