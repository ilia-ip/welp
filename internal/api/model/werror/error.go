package werror

import "net/http"

type Error struct {
	Code int
	Msg  string
}

var (
	BAD_REQ = Error{
		Code: http.StatusBadRequest,
		Msg:  "",
	}

	INT_ERR = Error{
		Code: http.StatusInternalServerError,
		Msg:  "Internal error",
	}
)
