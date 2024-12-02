package main

import (
	"fmt"
	"log"
	"onse/internal/email"
	"onse/internal/helpers"
)

func main() {
	config := helpers.NewConfig()
	smtp := config.Smtp

	if config.DebugMode {
		fmt.Printf("config: %+v\n\n", config)
	}

	participants, err := helpers.GetParticipants(config.CsvFile)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	fmt.Println("this year participants are:")
	for _, p := range participants {
		fmt.Printf("%d: %s (%s) with exclusion for %v\n", p.Id, p.Name, p.Email, p.Exclusions)
	}

	pairings := helpers.GeneratePairings(participants)
	emailDataList := helpers.GetEmailData(participants, pairings)

	for _, ed := range emailDataList {
		body, err := email.ParseHtml(config.EmailTemplate, ed)
		if err != nil {
			log.Fatalf("error parsing template: %s", err)
		}

		var receiverEmail string
		if config.DebugMode {
			receiverEmail = config.EmailReceiverTest
		} else {
			receiverEmail = ed.GifterEmail
		}

		smtpErr := smtp.Send(config.EmailSender, receiverEmail, body)
		if smtpErr != nil {
			log.Fatalf("error sending email: %s", err)
		}
	}
	fmt.Println("\n📬 emails sent!")
}
