package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func main() {
	Mes := &Message {
		SuccessURL: "https://blog.csdn.net/huwh_/article/details/77140664?utm_medium=",
	}
	auth = smtp.PlainAuth("", "1136089132@qq.com", "utkuzwbwgdnaffaf", "smtp.qq.com")
	r := NewRequest([]string{"1351169665@qq.com"}, "hello,email!", "hello,world!")
	if err := r.TempOK("./templates/register-ok.html", Mes); err == nil {
		ok, _ := r.SendEmail()
		fmt.Println(ok)
	}
}

type Request struct{
	from    string
	to      []string
	subject string
	body    string
}

type Message struct {
	ActiveURL  string
	UserMes    struct {
		UserName string
		Phone    string
		Email    string
		Company  string
		CheckURL string
	}
	SuccessURL   string
	FailURL      string
	Captcha      string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request {
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + r.body)
	addr :=  "smtp.qq.com:25"

	if err := smtp.SendMail(addr, auth, "1136089132@qq.com", r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) TempOK(templatePathName string, Mes *Message) error {
	fmt.Println("aaa")
	t, err := template.ParseFiles(templatePathName)
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.Execute(buf, Mes); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}