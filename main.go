package main

import (
	"cli-notes/greetings"
	"errors"
	"fmt"
	"log"
)

var id uint

type Record struct {
	Note string
	id   uint
}

func main() {

	log.SetFlags(0)

	var userInput string

	var notes = make([]Record, 0)

	for {

		greetings.PrintHello()

		fmt.Scanln(&userInput)

		switch userInput {
		case "/l":
			for _, note := range notes {
				fmt.Printf("note id: %v\nnote: %v\n", note.id, note.Note)
				fmt.Println("----------")
			}
		case "/q":
			break
		default:
			note, err := createRecord(userInput)

			if err != nil {
				log.Println(err)
				continue
			}

			notes = append(notes, note)

		}
	}
}

func createRecord(userInput string) (Record, error) {

	if userInput == "" {
		return Record{}, errors.New("empty record")
	}

	id = id + 1

	newRecord := Record{
		Note: userInput,
		id:   id,
	}

	return newRecord, nil

}
