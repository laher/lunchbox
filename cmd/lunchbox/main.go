package main

import (
	"fmt"
	"os"

	"github.com/laher/lunchbox/lunch"
)

func main() {
	subcommand := ""
	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}
	switch subcommand {
	case "script":
		err := lunch.ElvishRunScript(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "shell":
		err := lunch.ElvishPrompt(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "date":
		err := lunch.Date(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "http":
		err := lunch.HTTP(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "json": // deprecated .. use jq instead
		err := lunch.JSON(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "jq":
		err := lunch.JQ(os.Args[2:])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	// TODO:
	// case "cp":
	// case "mv":
	// case "rm":
	// case "http":
	// case "uncompress":
	// case "compress":

	default:
		fmt.Println("expected a subcommand")
		os.Exit(1)
	}
}
