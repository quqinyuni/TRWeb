package tool

import (
	"github.com/matcornic/hermes/v2"
	"strconv"
	"strings"
	"unsafe"
)
func convert( b []byte ) string {
	s := make([]string,len(b))
	for i := range b {
		s[i] = strconv.Itoa(int(b[i]))
	}
	return strings.Join(s,",")
}
func  String(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
func CreatedMail(Username string,CheckUrl string)(string){
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		// Theme: new(Default)
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: "OWUHU",
			Link: "https://www.owuhu.com/",
			// Optional product logo
			//Logo: "http://www.duchess-france.org/wp-content/uploads/2016/01/gopher.png",
			Copyright: "Copyright © 2021 Dharma Initiative. All rights reserved.",
			TroubleText: "如果您的浏览器不支持按钮点击，请复制下面的链接粘贴到网址栏打开。",
		},
	}

	email := hermes.Email{
		Body: hermes.Body{
			Name: Username,
			Signature: "你的朋友",
			Intros: []string{
				"欢迎来到 OWUHU ，非常欢迎你的加入。",
			},
			Actions: []hermes.Action{
				{
					Instructions: "请点击按钮，对您的账户进行验证。",
					Button: hermes.Button{
						Color: "#22BC66", // Optional action button color
						Text:  "点击验证",
						Link:  CheckUrl,
					},
				},
			},
			Outros: []string{
				"需要帮助或有问题，请回复此邮件，我们会在24小时内对您进行答复。",
			},
		},
	}

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(email)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	//emailText, err := h.GeneratePlainText(email)
	//if err != nil {
	//	panic(err) // Tip: Handle error with something else than a panic ;)
	//}

	// Optionally, preview the generated HTML e-mail by writing it to a local file

	return String([]byte(emailBody))
}