package notify

import (
	"fmt"
	"log"
	"net/url"

	"github.com/jhidalgo3/aremysiteup/params"
	mailgun "github.com/mailgun/mailgun-go"
)

// Mailer struct
type Mailer struct {
	To, From string
	Mailgun  params.Mailgun
}

// BuildAndSendEmail send alert email
func (m *Mailer) ComposeAndSendEmail(errors []error) {
	if m.To != "" {
		var (
			count   = len(errors)
			subject string
			body    string
		)
		if count == 1 {
			urlErr, ok := errors[0].(*url.Error)
			if ok {
				subject = fmt.Sprintf("'%s' is unreachable", urlErr.URL)
			} else {
				subject = "A site is unreachable"
			}
		} else {
			subject = fmt.Sprintf("%d sites are unreachable", count)
		}
		for _, error := range errors {
			body += fmt.Sprintf("* %s\n\n", error.Error())
		}
		err := m.sendEmail(subject, body)
		if err != nil {
			log.Println(err)
		}
	}
}

func (m *Mailer) sendEmail(subject, body string) error {
	log.Printf("sendEmail %s \n %s %s", m.To, subject, body)
	msg := mailgun.NewMessage(
		m.From,
		"[Aremysiteup] "+subject,
		body,
		m.To)

	mg := mailgun.NewMailgun(m.Mailgun.Domain, m.Mailgun.ApiKey, m.Mailgun.PublicApiKey)

	resp, id, err := mg.Send(msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Printf("ID: %s Resp: %s\n", id, resp)

	return nil
}
