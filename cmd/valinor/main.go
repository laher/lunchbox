package main

import (
	"fmt"
	"os"

	"github.com/laher/valinor.elv/valinor"
)

func main() {
	subcommand := ""
	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}
	switch subcommand {
	case "script":
		err := valinor.ElvishRunScript(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "shell":
		err := valinor.ElvishPrompt(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "date":
		err := valinor.Date(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("expected a subcommand")
		os.Exit(1)
	}
}
