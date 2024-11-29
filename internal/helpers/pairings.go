package helpers

import (
	"fmt"
	"math/rand/v2"
	"onse/internal/email"
)

func GeneratePairings(participants []Participant) map[int]int {
	fmt.Println("\n🎁 drawing Secret Santa...")
	var pairings map[int]int
	var err error

	// retry until a valid pairing is found
	for {
		pairings, err = attemptPairings(participants)
		if err == nil {
			break
		}
		fmt.Println("🔁 retrying:", err)
	}

	fmt.Printf("🥁 ready!\n\n")
	return pairings
}

func attemptPairings(participants []Participant) (map[int]int, error) {
	shuffled := shuffle(participants)

	pairings := make(map[int]int)
	paired := make(map[int]bool)

	for i := 0; i < len(shuffled); i++ {
		gifter := shuffled[i]

		for j := 0; j < len(shuffled); j++ {
			receiver := shuffled[j]

			// do not pair in these cases:
			// - with themselves
			// - if the receiver is already paired with someone else
			// - if the receiver is in the exclusion list of the gifter
			// - if the pairing creates a bidirectional association
			if gifter.Id == receiver.Id || paired[receiver.Id] || find(gifter.Exclusions, receiver.Id) || pairings[receiver.Id] == gifter.Id {
				continue
			}

			pairings[gifter.Id] = receiver.Id
			paired[receiver.Id] = true
			break
		}
	}

	// check for any unmatched participants
	for _, p := range shuffled {
		if !paired[p.Id] {
			return nil, fmt.Errorf("participant %d was not paired", p.Id)
		}
	}

	return pairings, nil
}

// randomly reorder list
func shuffle(list []Participant) []Participant {
	fmt.Println("🎲 shuffling participants...")

	shuffled := append([]Participant(nil), list...)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	// for _, p := range shuffled {
	// 	fmt.Printf("%d - %s\n", p.Id, p.Name)
	// }
	// fmt.Printf("\n")

	return shuffled
}

// finds a value in a list
func find(list []int, value int) bool {
	for _, element := range list {
		if element == value {
			return true
		}
	}
	return false
}

func GetParticipantEmailData(participants []Participant, pairings map[int]int) []email.EmailData {
	participantsMap := make(map[int]Participant)
	var list []email.EmailData

	for _, p := range participants {
		participantsMap[p.Id] = p
	}

	for _, gifter := range participants {
		receiverId := pairings[gifter.Id]
		receiver := participantsMap[receiverId]

		list = append(list, email.EmailData{
			GifterName:   gifter.Name,
			GifterEmail:  gifter.Email,
			ReceiverName: receiver.Name,
			Message:      gifter.Message,
		})
	}

	return list
}
