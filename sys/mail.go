package sys

import (
	"net/smtp"
	"strings"
)

func SendMail(from string, to string, subject string, body string) error {
	user := "jiangyong_uxin@aliyun.com"
	password := "uxin27jiang"
	host := "smtp.aliyun.com:25"

	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	content_type = "Content-type:text/html;charset=utf-8"

	msg := []byte("To: " + to + "\r\nFrom: " + from + "<" + from + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}
