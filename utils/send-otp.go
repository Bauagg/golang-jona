package utils

import (
	"backend-jona-golang/config"
	"fmt"
	"net/smtp"
)

func SendEmail(to string, otp uint64) error {
	from := config.GMAIL_OTP
	password_otp := config.PASSWORD_OTP

	// SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	subject := "Your OTP Code"
	body := fmt.Sprintf("Your OTP code is: %d", otp)
	message := fmt.Sprintf("Subject: %s\n\n%s", subject, body)

	// Authentication.
	auth := smtp.PlainAuth("", from, password_otp, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(message))
	if err != nil {
		return err
	}

	return nil
}
