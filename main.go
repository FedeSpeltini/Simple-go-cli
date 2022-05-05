package main

import (
	"flag"
	"fmt"
	"strconv"

	"github.com/FedeSpeltini/go-cli/commands"
)

func main() {
	var expenses []float32
	var export string
	flag.StringVar(&export, "export", "", "export to file")

	flag.Parse()

	for {
		input, err := commands.GetInput()

		if err != nil {
			panic(err)
		}

		if input == "exit" {
			break
		}

		expense, err := strconv.ParseFloat(input, 32)
		if err != nil {
			panic(err)
		}
		expenses = append(expenses, float32(expense))

	}
	if export == "" {
		commands.ShowInConsole(expenses)
	} else {
		commands.Export(export, expenses)
	}

	fmt.Println("Bye!")
}
