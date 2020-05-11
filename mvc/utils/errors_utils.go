package utils

type ApplicationError struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
	Code   string `json:"code"`
}
