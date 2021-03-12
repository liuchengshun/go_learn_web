package main

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"
)

var auth smtp.Auth

func main() {
	Msg := &Message {
		ActiveURL:      "https://www.tfcloud.com/",
		UserMsg:         UserMsg{
			UserName:   "ZhangSan",
			Phone:      "123456",
			Email:      "132@qq.com",
			Company:    "tongfangyouyun",
			CheckURL:   "https://www.tfcloud.com/",
		},
		SuccessURL:     "https://www.tfcloud.com/",
		FailURL:        "https://www.tfcloud.com/",
		Captcha:        "7d8F",
	}
	// NewRequest(from string, to []string, subject, body, password, host string)
	r := NewRequest("1136089132@qq.com", []string{"1351169665@qq.com"}, "hello,email!", 
			"hello,world!", "rwkiovjiuqsihbhf", "smtp.qq.com")
	// if err := r.RenderTemplate("./templates/success.html", Msg); err == nil {
	// 	if ok, _ := r.SendEmail(); ok {
	// 		fmt.Println("register-ok sended")
	// 	}
	// }
	// if err := r.RenderTemplate("./templates/apply.html", Msg); err == nil {
	// 	fmt.Println("aaa")
	// 	if ok, _ := r.SendEmail(); ok {
	// 		fmt.Println("register-email sended")
	// 	}
	// }
	// if err := r.RenderTemplate("./templates/admin.html", Msg); err == nil {
	// 	if ok, _ := r.SendEmail(); ok {
	// 		fmt.Println("register-admin sended")
	// 	}
	// }
	// if err := r.RenderTemplate("./templates/failed.html", Msg); err == nil {
	// 	if ok, _ := r.SendEmail(); ok {
	// 		fmt.Println("register-failed sended")
	// 	}
	// }
	if err := r.RenderTemplate("./templates/captcha.html", Msg); err == nil {
		if ok, _ := r.SendEmail(); ok {
			fmt.Println("register-captcha sended")
		}
	}
}

type Request struct{
	to       []string
	subject  string
	body     string
	from     string
	password string
	host     string
}

type UserMsg struct {
	UserName     string
	Phone        string
	Email        string
	Company      string
	CheckURL     string
}

type Message struct {
	ActiveURL    string
	UserMsg      UserMsg
	SuccessURL   string
	FailURL      string
	Captcha      string
}

func NewRequest(from string, to []string, subject, body, password, host string) *Request {
	return &Request {
		to:       to,
		subject:  subject,
		body:     body,
		from:     from,
		password: password,
		host:     host,
	}
}

func (r *Request) Auth() smtp.Auth {
	return smtp.PlainAuth("", r.from, r.password, r.host)
}

func (r *Request) SendEmail() (bool, error) {
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + r.body)
	auth := r.Auth()

	if err := smtp.SendMail(r.host + ":25", auth, r.from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) RenderTemplate(templatePathName string, Mes *Message) error {
	t, err := template.ParseFiles(templatePathName, "./static/header.html", "./static/footer.html")
	if err != nil {
		return err
	}

	buf := new(bytes.Buffer)
	if err = t.ExecuteTemplate(buf, "content", Mes); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}
