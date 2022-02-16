package mail

import (
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/model"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"net/smtp"
	"os"
	"strings"
)

const (
	HOST    = "smtp.yandex.ru"
	PORT    = "587"
	SUBJECT = "Food Delivery"
)

var Post = make(chan model.Post, 1)

func SendEmail(logger logging.Logger, auth smtp.Auth) {
	for {
		select {
		case <-Post:
			destination := <-Post
			from := os.Getenv("POST_FROM")
			smtpHost := HOST
			smtpPort := PORT
			msg := fmt.Sprintf(" Уважаемый клиент, Ваш текущий пароль: %s.", destination.Password)
			message := strings.Replace("From: "+from+"~To: "+destination.Email+"~Subject: "+SUBJECT+"~~", "~", "\r\n", -1) + msg
			err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{destination.Email}, []byte(message))
			if err != nil {
				logger.Errorf("Error while sending email to %s:%s", destination.Email, err)
				return
			}
			logger.Infof("Email for %s Sent Successfully!", destination.Email)
		}

	}
}
