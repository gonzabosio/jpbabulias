package async

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/gonzabosio/jpbabulias/db/model"
	"github.com/jordan-wright/email"
)

func SendApptEmail(appt *model.InsertAppointment) {
	go func() {
		bsnEmail := os.Getenv("BUSINESS_EMAIL")
		em := email.Email{
			From:    fmt.Sprintf("Od. Babulias <%v>", bsnEmail),
			To:      []string{appt.UserEmail},
			Subject: "Turno agendado con dentista",
			Text: []byte(fmt.Sprintf("Turno agendado con el odontólogo Babulias para %s el día: %v\nAsunto: %s",
				appt.FullName,
				appt.ApptDate.Format("02-01-2006 15:04"),
				appt.Subject),
			),
		}
		if err := em.Send("smtp.gmail.com:587", smtp.PlainAuth("", bsnEmail, os.Getenv("BUSINESS_EMAIL_PW"), "smtp.gmail.com")); err != nil {
			log.Println("EMAIL ERROR: Email sending failed ")
		}
	}()
}
