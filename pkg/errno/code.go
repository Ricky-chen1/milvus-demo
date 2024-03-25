package errno

const (
	SuccessCode      = 200
	SuccessMsg       = "ok"
	ServiceErrorCode = 500
)

var (
	Success      = NewErrNo(SuccessCode, SuccessMsg)
	ServiceError = NewErrNo(ServiceErrorCode, "Server error")
)
