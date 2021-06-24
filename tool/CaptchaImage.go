package tool

import (
	"github.com/mojocn/base64Captcha"
	"tongren/e"
)

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

func GenerateCaptcha() (string, string) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 8)
	captcha := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := captcha.Generate()
	if err != nil {
		id, b64s = "error", "error"
	}
	return id, b64s
}

// base64Captcha verify http handler
func CaptchaVerify(Id string, VerifyValue string) (int) {
	code := e.SUCCESS
	if !store.Verify(Id, VerifyValue, true) {
		code = e.ERROR_CAPT_VALUE
	}
	return code
}
