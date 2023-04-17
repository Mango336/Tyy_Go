package errmsg

const (
	SUCCESS = 200
	ERROR   = 500

	// 登录错误 100x
	ERROR_USER_NOT_EXIST = 1001
)

var errMp = map[int]string{
	SUCCESS: "OK",
	ERROR:   "Error",

	ERROR_USER_NOT_EXIST: "用户名或密码错误...",
}

func GetErrMsg(i int) string {
	return errMp[i]
}
