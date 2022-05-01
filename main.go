package main

import "fmt"

var id = 0

func main() {

	var userInput string
	type Record struct {
		Note string
		id   int
	}
	var notes = make([]Record, 0)

	for {
		fmt.Printf("Enter your note:\n")
		fmt.Scan(&userInput)

		switch userInput {
		case "/l":
			for _, note := range notes {
				fmt.Printf("note id: %v\nnote: %v\n", note.id, note.Note)
				fmt.Println("----------")
			}
		case "/q":
			break
		default:
			id = id + 1

			newRecord := Record{
				Note: userInput,
				id:   id,
			}

			notes = append(notes, newRecord)
		}

	}

}
