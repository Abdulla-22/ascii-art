package ascii_art

import (
	"bufio"
	// "fmt"
	"log"
	"os"
)

var PrintStr string

func TakeFromFile(LineNumber int, BannaerName string) {

	file, err := os.Open(TakeBanner(BannaerName))

	if err != nil {
		log.Fatal("Error to opne the standard.txt file")

	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	counter := 0

	for scanner.Scan() {
		counter++

		if counter == LineNumber {
			PrintStr = scanner.Text()
			break
		}
	}

	//fmt.Print(PrintStr)
	result = result + PrintStr
}
