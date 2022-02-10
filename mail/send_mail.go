package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

const (
	HOST        = "smtp.gmail.com"
	PORT        = "587"
	USER        = ""
	PASSWORD    = ""
)

func SendEmail(to, subject, msg string) error {
	auth := smtp.PlainAuth("", USER, PASSWORD, HOST)
	addr := fmt.Sprintf("%s:%s", HOST, PORT)
		{
			str := strings.Replace("From: "+USER+"~To: "+to+"~Subject: "+subject+"~~", "~", "\r\n", -1) + msg
			err := smtp.SendMail(addr, auth, USER,	[]string{to}, []byte(str))
			if err != nil {
				log.Println(err)
				os.Exit(1)
		}
		log.Println("Successfully sent mail")
		return nil
	}
}