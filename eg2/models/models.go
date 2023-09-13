package models

type Sample struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type DBResponse struct {
	Username string `json:"username"`
}
