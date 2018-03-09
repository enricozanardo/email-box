package email

import (
	"net/smtp"
	"bytes"
	"html/template"
	"github.com/spf13/viper"
	"github.com/goinggo/tracelog"
	"os"
)

var auth smtp.Auth

//Request struct
type Request struct {
	from    string
	to      []string
	subject string
	body    string
}

func NewRequest(to []string, subject, body string) *Request {
	return &Request{
		to:      to,
		subject: subject,
		body:    body,
	}
}

func (r *Request) SendEmail() (bool, error) {

	username := viper.GetString("email.username")
	password := viper.GetString("email.password")
	host := viper.GetString("email.host")
	from := viper.GetString("email.from")

	auth = smtp.PlainAuth("", username, password , host)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject := "Subject: " + r.subject + "!\n"
	msg := []byte(subject + mime + "\n" + r.body)
	addr := viper.GetString("email.addr")

	if err := smtp.SendMail(addr, auth, from, r.to, msg); err != nil {
		return false, err
	}
	return true, nil
}

func (r *Request) ParseTemplate(templateFileName string, data interface{}) error {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return err
	}
	r.body = buf.String()
	return nil
}


func SendConfirmRegistrationEmail(email string, token string) bool {

	viper.SetConfigName("config")
	viper.AddConfigPath("../")

	if err := viper.ReadInConfig(); err != nil {
		tracelog.Errorf(err, "email", "SendConfirmRegistrationEmail", "Error reading config file")
	}

	// Set in Rancher a variable URL http://172.104.230.81/confirm?k=
	url := os.Getenv("URL")

	if len(url) == 0 {
		//use local variables
		url = viper.GetString("service.urlLocal")
	}

	templateData := struct {
		Email string
		URL  string
	}{
		Email: email,
		URL:  url + token,
	}

	textBody := "Thank you! We have received a request from this email account: " + email + " for " +
		"registering to the OneZero Binary website. In order to enable your account you have to confirm " +
		"it clicking on the following link: " + url + " within the 23:59 of the current day." +
		"Hope to see you soon! OneZero Binary Team"

	r := NewRequest([]string{email}, "Confirm Registration", textBody)

	err := r.ParseTemplate("./template.html", templateData)

	if err != nil {
		tracelog.Errorf(err, "email", "SendConfirmRegistrationEmail", "Error parsing the template")
	}

	ok, _ := r.SendEmail()

	return ok
}

