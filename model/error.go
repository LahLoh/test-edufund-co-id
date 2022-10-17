package model

type ErrorResponse struct {
	Msg  string `json:"error"`
	Code int    `json:"-"`
}

func (er ErrorResponse) Error() string {
	return er.Msg
}
