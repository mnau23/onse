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
	sender := config.EmailSender
	template := config.EmailTemplate
	ParticipantsList := config.CsvFile

	if config.DebugMode {
		fmt.Printf("config is %+v\n\n", config)
	}

	// read and parse the CSV file
	participants, err := helpers.GetParticipants(ParticipantsList)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	for _, p := range participants {
		fmt.Printf("Participant %s has email %s\n", p.Name, p.Email)
	}

	// TODO: setup email data for each participant
	data := email.EmailData{
		Name:        "John Doe",
		SecretSanta: "Jane Doe",
		Message:     "a random text here",
	}
	// TODO: then add it to template
	body, err := email.ParseTemplate(template, data)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
	// fmt.Printf("\nBody: %s\n", body)

	// TODO: finally send email
	err = smtp.Send(sender, "todo-participant-email", body)
	if err != nil {
		fmt.Println("error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
