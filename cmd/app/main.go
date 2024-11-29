package main

import (
	"fmt"
	"log"
	"onse/internal/email"
	"onse/internal/helpers"
)

func main() {
	// load app configuration
	config := helpers.NewConfig()
	smtp := config.Smtp
	emailSender := config.EmailSender
	emailTemplate := config.EmailTemplate
	participantsList := config.CsvFile

	if config.DebugMode {
		fmt.Printf("config is %+v\n\n", config)
	}

	participants, err := helpers.GetParticipants(participantsList)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	fmt.Println("this year participants are:")
	for _, p := range participants {
		fmt.Printf("%d: %s - %s with exclusions on %v\n", p.Id, p.Name, p.Email, p.Exclusions)
	}

	pairings := helpers.GeneratePairings(participants)
	for gifter, receiver := range pairings {
		fmt.Printf("Participant %d has %d\n", gifter, receiver)
	}
	fmt.Print("\n")

	emailDataList := helpers.GetParticipantEmailData(participants, pairings)

	for _, ed := range emailDataList {
		body, err := email.ParseHtml(emailTemplate, ed)
		if err != nil {
			log.Fatalf("error parsing template: %s", err)
		}
		smtpErr := smtp.Send(emailSender, "todo-participant-email", body)
		if smtpErr != nil {
			log.Fatalf("error sending email: %s", err)
		}
	}
	fmt.Println("\n📬 emails sent!")
}
