package constant

const (
	SUCCSE = 200
	ERROR  = 500

	// code= 1000... 用户模块的错误
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NO_RIGHT    = 1008
	// code= 2000... 视频模块的错误

	ERROR_ART_NOT_EXIST = 2001
	// code= 3000... 视频模块的错误
	ERROR_CATENAME_USED  = 3001
	ERROR_CATE_NOT_EXIST = 3002
	//code = 4000... 用户follow模块错误
	ERROR_Follow_NOT_EXIST = 4001
)

var codeMsg = map[int]string{
	SUCCSE:                 "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME_USED:    "用户名已存在！",
	ERROR_PASSWORD_WRONG:   "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在,请重新登陆",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期,请重新登陆",
	ERROR_TOKEN_WRONG:      "TOKEN不正确,请重新登陆",
	ERROR_TOKEN_TYPE_WRONG: "TOKEN格式错误,请重新登陆",
	ERROR_USER_NO_RIGHT:    "该用户无权限",

	ERROR_ART_NOT_EXIST: "视频不存在",

	ERROR_CATENAME_USED:  "该标签已存在",
	ERROR_CATE_NOT_EXIST: "该标签不存在",

	ERROR_Follow_NOT_EXIST: "未关注",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
