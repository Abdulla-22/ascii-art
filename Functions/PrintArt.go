package ascii_art

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var result string

func PrintArt(str, BannaerName string) {
	result = ""
	if str == "\\n" {
		result = result + "\n"
	} else {
		Sentence := strings.Split(str, "\\n")

		for i := 0; i < len(Sentence); i++ {

			if Sentence[i] == "" {

				if !(i+1 == len(Sentence)) || !(Sentence[(len(Sentence)-1)] == "") {
					result = result + "\n"
				}

			} else {
				// Make a function to go throw the leter line by line.
				PrintLintByLine(Sentence[i], BannaerName)

			}

		}
	}
	if Output != "" {
		// TODO: create the file and write the result there.

		if strings.HasSuffix(Output, ".txt") {
		} else {
			Output = Output + ".txt"
		}

		err := os.WriteFile(Output, []byte(result), 0666)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Print(result)
	}
}
