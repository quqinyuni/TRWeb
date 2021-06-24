package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"tongren/e"
	"tongren/tool"
)

type UserForm struct {
	User        string `form:"user" binding:"required"`
	Password    string `form:"password" binding:"required"`
	Email string `form:"email" binding:"required"`
	Id          string `form:"captid" binding:"required"`
	VerifyValue string `form:"verifyvalue" binding:"required"`
}
//定义变量
var form UserForm
/*注册*/
func Register(r *gin.Context) {
	//绑定格式
	r.ShouldBindWith(&form, binding.Form)
	//默认200，错误返回对应码
	code := tool.CaptchaVerify(form.Id, form.VerifyValue)

	err := tool.SendToMail(
		"quqinyuni",
		form.Email,
		"邮箱激活",
		tool.CreatedMail("test",tool.CreatCheckUrl(form.Id)))
	if err != nil {
		code = e.ERROR_MAIL_SEND
	}
	r.JSON(http.StatusOK, gin.H{
		"message": e.GetMsg(code),
	})
}

/*登录*/
func Signin(r *gin.Context) {
	//绑定格式
	r.ShouldBindWith(&form, binding.Form)
	code := tool.CaptchaVerify(form.Id, form.VerifyValue)
	/*
		TODO
		验证账号密码及验证码是否正确
	*/
	r.JSON(http.StatusOK, gin.H{
		"message": e.GetMsg(code),
	})
}
