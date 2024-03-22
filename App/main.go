package main

import (
	ascii_art_output "ascii_art_output/Functions"
	"log"
	"os"
	"strings"
)

func main() {
	The_Args := os.Args
	//Banner := "standard"
	if len(The_Args) == 1 {
		log.Fatal("You have entertained no arguments. \nPlease insert your argument between double quotation like the following: \n$ go run . \"your argument\"\n")

	} else if len(The_Args) == 2 {
		if The_Args[1] == "" {
			return
		} else {
			// Cheek the text if it contains only English and printable character from the ASCII table
			for _, ascii := range The_Args[1] {
				if ((ascii < 32) || (ascii > 126)) && (ascii != 10) {
					log.Fatal("USE ONLY ENGLISH")
					return
				}
			}
			// Print the art with the stander format and white colore
			ascii_art_output.PrintArt(The_Args[1], "")

		}

		//Handl 3 arg (banars name)
	} else if len(The_Args) == 3 {
		if The_Args[1] == "" {
			return

		} else {
			// Cheek the text if it contains only English and printable character from the ASCII table
			word := The_Args[1]
			if strings.HasPrefix(The_Args[1], "--output=") {
				ascii_art_output.Output = strings.TrimPrefix(The_Args[1], "--output=")
				The_Args[1] = The_Args[2]
				The_Args[2] = ""
			}

			for _, ascii := range word {
				if ((ascii < 32) || (ascii > 126)) && (ascii != 10) {
					log.Fatal("USE ONLY ENGLISH")
					return
				}
			}

			// Print the art with the stander format and white colore
			ascii_art_output.PrintArt(The_Args[1], The_Args[2])

		}
	} else if len(The_Args) == 4 {
		if The_Args[1] == "" {
			return

		} else {
			// Cheek the text if it contains only English and printable character from the ASCII table
			word := The_Args[2]
			if strings.HasPrefix(The_Args[1], "--output=") {
				ascii_art_output.Output = strings.TrimPrefix(The_Args[1], "--output=")
			} else {
				log.Fatal("You have entertained too many arguments. \nPlease insert your argument between double quotation like the following: \n$ go run . \"your argument\"\n")
			}
			for _, ascii := range word {
				if ((ascii < 32) || (ascii > 126)) && (ascii != 10) {
					log.Fatal("USE ONLY ENGLISH")
					return
				}
			}
			// Print the art with the stander format and white colore
			ascii_art_output.PrintArt(The_Args[2], The_Args[3])

		}
	} else if len(The_Args) > 4 {
		log.Fatal("You have entertained too many arguments. \nPlease insert your argument between double quotation like the following: \n$ go run . \"your argument\"\n")
	}

}
