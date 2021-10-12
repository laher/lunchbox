package main

import (
	"context"
	"fmt"
	"os"

	"github.com/laher/lunchbox/lunch"
)

func main() {
	subcommand := ""
	if len(os.Args) > 1 {
		subcommand = os.Args[1]
	}
	ctx := lunch.Context{
		Ctx: context.Background(),
	}
	lun, ok := lunch.Get(subcommand)

	if !ok {
		fmt.Println("expected a valid subcommand")
		os.Exit(1)
	}
	if err := lun(ctx, os.Args[2:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO:
	// case "cp":
	// case "mv":
	// case "rm":
	// case "uncompress":
	// case "compress":

}
