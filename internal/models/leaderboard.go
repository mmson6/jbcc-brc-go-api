package models

type Leaderboard struct {
	Users []User `json:"users"`
}

type User struct {
	DisplayName string `json:"displayName"`
	VerseCount  int    `json:"verseCount"`
}
