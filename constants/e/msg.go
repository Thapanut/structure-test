package e

var MsgFlags = map[int]string{
	ERROR:                 "error",
	SUCCESS:               "success",
	CREATED:               "created",
	BAD_REQUEST:           "bad request",
	ERROR_AUTH:            "you don't have permission to access",
	ERROR_FORBIDDEN:       "forbidden",
	DATA_NOT_FOUND:        "data not found",
	INVALID_PARAMS:        "invalid params",
	CONFLICT:              "conflict",
	INTERNAL_SERVER_ERROR: "internal server error",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
