package email

type EmailSender struct {
	SmtpHost string `env:"SMTP_HOST,required"`
	SmtpPort int    `env:"SMTP_PORT,required"`
	Username string `env:"SMTP_USER,required"`
	Password string `env:"SMTP_PASS,required"`
	From     string `env:"EMAIL_FROM,required"`
}
