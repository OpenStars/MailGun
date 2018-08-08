package mailgun

import (
	"net/smtp"

	"github.com/OpenStars/email"
)

type MailGun struct {
	AuthID  string
	AuthPwd string
}

func onceInitMailGunInstance(authID, authPwd string) {
	mailGunInstance = NewMailGun(authID, authPwd)
}

func NewMailGun(authID, authPwd string) *MailGun {
	return &MailGun{
		AuthID: authID,
		AuthPwd: authPwd,
	}
}

func (m *MailGun) SendMail(receiver []string, subject, content string) {
	e := email.NewEmail()
	e.From = m.AuthID
	e.To = receiver
	e.Subject = subject
	//e.Text = []byte("Thử mailgun xem có ra cái khỉ gì không")
	e.HTML = []byte(content)

	go func(mail *email.Email, authID, authPwd string) {
		mail.Send("smtp.mailgun.org:587", smtp.PlainAuth("", authID, authPwd, "smtp.mailgun.org"))
	}(e, m.AuthID, m.AuthPwd)
}