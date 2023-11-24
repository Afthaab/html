package pkg

import (
	"crypto/rand"
	"errors"
	"fmt"
	"net/smtp"
)

func GenCaptchaCode() (string, error) {
	codes := make([]byte, 6)
	if _, err := rand.Read(codes); err != nil {
		return "", err
	}

	for i := 0; i < 6; i++ {
		codes[i] = uint8(48 + (codes[i] % 10))
	}

	return string(codes), nil
}

func GenerateOneTimePassword(gmail string) (string, error) {
	// Sender's email address and password
	from := "footballestore@gmail.com"
	password := "dter dlkf dwsb vmgt"

	// Recipient's email address
	to := gmail

	// SMTP server details
	smtpServer := "smtp.gmail.com"
	smtpPort := 587

	otp, err := GenCaptchaCode()
	if err != nil {
		return "", errors.New("error in creating the ramdom number")
	}

	// Message content
	message := []byte("your one time password " + otp)

	// Authentication information
	auth := smtp.PlainAuth("", from, password, smtpServer)

	// SMTP connection
	smtpAddr := fmt.Sprintf("%s:%d", smtpServer, smtpPort)
	err = smtp.SendMail(smtpAddr, auth, from, []string{to}, message)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return "", err
	}

	fmt.Println("Email sent successfully!")
	return otp, nil
}
