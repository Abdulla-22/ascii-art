package ascii_art

func PrintLintByLine(str, BannaerName string) {

	for row := 0; row < 8; row++ {

		for columns := 0; columns < len(str); columns++ {
			character := str[columns]

			TakeFromFile((Line_num(character) + row), BannaerName)
		}
		result = result + "\n"
	}
	//Take new line after printing the art

}
