package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tongren/e"
	"tongren/tool"
)
type Capt struct {
	CaptId          string `form:"captid" json:"captid" binding:"required"`
	CaptValue string `form:"captvalue" json:"captvalue" binding:"required"`
}
var captform Capt
func GetCapt(r *gin.Context){
	id,b64s := tool.GenerateCaptcha()
	r.JSON(http.StatusOK,gin.H{
		"id":id,
		"b64s":b64s,
	})
}

func CheckCapt(r *gin.Context)  {
	err := r.ShouldBindJSON(&captform)
	if err != nil{
		r.JSON(http.StatusOK,gin.H{
			"code":e.INVALID_PARAMS,
			"msg":e.GetMsg(e.INVALID_PARAMS),
			"error":err.Error(),
		})
		return
	}
	code := tool.CaptchaVerify(captform.CaptId,captform.CaptValue)
	r.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"form":captform,
	})
}
