package main

import (
	"fmt"
	"log"
	"onse/internal/email"
	"onse/internal/helpers"
)

func main() {
	// load app configuration
	Config := helpers.NewConfig()
	Smtp := Config.Smtp
	EmailSender := Config.EmailSender
	EmailTemplate := Config.EmailTemplate
	ParticipantsList := Config.CsvFile

	if Config.DebugMode {
		fmt.Printf("config is %+v\n\n", Config)
	}

	// read and parse the CSV file
	Participants, err := helpers.GetParticipants(ParticipantsList)
	if err != nil {
		log.Fatalf("error reading CSV: %s", err)
	}

	for _, p := range Participants {
		fmt.Printf("Participant %s has email %s\n", p.Name, p.Email)
	}

	// TODO: setup email data for each participant
	data := email.EmailData{
		Name:        "John Doe",
		SecretSanta: "Jane Doe",
		Message:     "a random text here",
	}
	// TODO: then add it to template
	body, err := email.ParseTemplate(EmailTemplate, data)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return
	}
	// fmt.Printf("\nBody: %s\n", body)

	// TODO: finally send email
	err = Smtp.Send(EmailSender, "todo-participant-email", body)
	if err != nil {
		fmt.Println("error sending email:", err)
		return
	}

	fmt.Println("Email sent successfully!")
}
