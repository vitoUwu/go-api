package utils

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

type APIError struct {
	Message string `json:"message"`
}
