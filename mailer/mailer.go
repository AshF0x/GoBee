package mailer

import (
	"fmt"
	"net/smtp"
)

// Mail ...
func Mail(updateList []byte) {
	// Sender data
	senderMail := ""
	senderPass := ""
	// Receiver Mail
	receiver := []string{
		"",
	}
	// SMTP Config
	smtpHost := ""
	smtpPort := ""
	message := []byte(updateList)
	auth := smtp.PlainAuth("", senderMail, senderPass, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, senderMail, receiver, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mail has been sent")
}
