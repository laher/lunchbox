package lunch

import (
	"errors"
	"flag"
	"fmt"

	xbarplugins "github.com/laher/lunchbox/xbar-plugins"
)

func init() {
	Register("xbar-plugin", XbarPlugin)
}

func XbarPlugin(ctx Context, args []string) error {
	var (
		cm = flag.NewFlagSet("xbar-plugin", flag.ExitOnError)
		ls = cm.Bool("list", false, "list available plugins")
	)
	if err := cm.Parse(args); err != nil {
		return err
	}

	if *ls {
		// list plugins
		ls, err := xbarplugins.List(ctx.Ctx)
		if err != nil {
			return err
		}
		for _, l := range ls {
			fmt.Println(l)
		}
		return nil
	}
	// install a plugin
	return errors.New("unimplemented")
}
