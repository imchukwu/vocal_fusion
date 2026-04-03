package email

import (
	"fmt"
	"net/smtp"
	"vocal_fusion/config"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type emailService struct {
	cfg *config.AppConfig
}

func NewEmailService(cfg *config.AppConfig) EmailService {
	return &emailService{cfg: cfg}
}

func (s *emailService) SendEmail(to string, subject string, body string) error {
	if s.cfg.SMTPHost == "" || s.cfg.SMTPPort == "" {
		return fmt.Errorf("SMTP configuration is incomplete")
	}

	auth := smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPass, s.cfg.SMTPHost)
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	msg := []byte(fmt.Sprintf("To: %s\r\nSubject: %s\r\n%s\r\n\r\n%s", to, subject, mime, body))

	addr := fmt.Sprintf("%s:%s", s.cfg.SMTPHost, s.cfg.SMTPPort)
	err := smtp.SendMail(addr, auth, s.cfg.SMTPFrom, []string{to}, msg)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	return nil
}
