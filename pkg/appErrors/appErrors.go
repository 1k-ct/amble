package appErrors

import "net/http"

type ApplicationError struct {
	Code  int    `json:"code"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
}

type ApplicationErrorMsg struct {
	Code     int    `json:"code"`
	Level    string `json:"level"`
	Msg      string `json:"msg"`
	ErrorMsg string `json:"error_msg"`
}

func ErrMeatdataMsg(err error, appErr *ApplicationError) *ApplicationErrorMsg {
	appError := &ApplicationErrorMsg{
		Code:     appErr.Code,
		Level:    appErr.Level,
		Msg:      appErr.Msg,
		ErrorMsg: err.Error(),
	}
	return appError
}

var ServerError = &ApplicationError{
	Code:  http.StatusInternalServerError,
	Level: "Error",
	Msg:   "An error has occurred inside the server.",
}
var ErrorJSON = &ApplicationError{
	Code:  http.StatusBadRequest,
	Level: "Error",
	Msg:   "I couldn't read the json.",
}
var ErrRecordDatabase = &ApplicationError{
	Code:  http.StatusBadRequest,
	Level: "Error",
	Msg:   "The ID does not exist. (Database error) 存在しないIDです。（データベースがエラー）",
}
var ErrNotCreateToken = &ApplicationError{
	Code:  http.StatusUnprocessableEntity,
	Level: "Error",
	Msg:   "This is a token production failure. token制作失敗です。",
}
var ErrInvalidToken = &ApplicationError{
	Code:  http.StatusUnauthorized,
	Level: "Error",
	Msg:   "The token is invalid. tokenが無効です。",
}
