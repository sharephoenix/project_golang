package email

import (
	"fmt"
	"log"
	"testing"
)

func Test_SendMail(t *testing.T) {
	//定义收件人
	mailTo := []string{
		"326083325@qq.com",
	}
	//邮件主题为"Hello"
	subject := "Hello by golang gomail from exmail.qq.com"
	// 邮件正文
	body := "Hello,by gomail sent"
	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err.Error())
		fmt.Println("send fail")
		return
	}

	fmt.Println("send successfully")
}
