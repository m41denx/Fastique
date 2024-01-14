package models

type ApiError struct {
	Message string `json:"message"`
}

type ApiLoginResponse struct {
	Token string `json:"token"`
}
