package e

var MsgFlags = map[int]string{
	SUCCESS:        "SUCCESS",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_AUTH:                     "用户登录错误",
	ERROR_AUTH_TOKEN:               "token获取失败",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "用户token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "用户token鉴权超时",

	ERROR_CUSTOMER:        "客户信息不存在",
	ERROR_CUSTOMER_INSET:  "客户信息写入错误",
	ERROR_CUSTOMER_EDIT:   "客户信息修改失败",
	ERROR_CUSTOMER_DELETE: "客户信息删除失败",

	ERROR_CAPT_VALUE: "验证码输入错误",
	ERROR_MAIL_SEND:"邮件发送失败",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}