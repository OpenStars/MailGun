package mailgun

import "sync"

var (
	once sync.Once

	mailGunInstance    MailGunIf
)

func GetMailGun(authID, authPwd string) MailGunIf {
	once.Do(func() {
		onceInitMailGunInstance(authID, authPwd)
	})
	return mailGunInstance
}


type MailGunIf interface {
	SendMail(receiver []string, subject, content string)
}