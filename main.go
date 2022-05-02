package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	id3 "github.com/mikkyang/id3-go"
)

func AskConfirmation() bool {
	var s string

	fmt.Printf("(y/N): ")
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}

	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	if s == "y" || s == "yes" {
		return true
	}
	return false
}

func main() {
	mp3 := flag.String("mp3", "", "input mp3 file.")
	flag.Parse()

	fmt.Println("This is a test of the ID3 tag module.")
	file, err := id3.Open(*mp3)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("The current Album for the mp3 is: \"" + file.Artist() + "\". Would you like to change it?")
	isConfirmed := AskConfirmation()
	if isConfirmed {
		fmt.Println("What would you like to change? (artist): ")
		var change string
		change = strings.ToLower(change)
		fmt.Scanln(&change)

		fmt.Println("What would you like to change it to? ")
		var changeto string
		fmt.Scanln(&changeto)

		if change == strings.ToLower("Artist") {
			file.SetArtist(changeto)
		}

	} else {
		fmt.Println("Choose the right file then, asshole!")
	}

}
