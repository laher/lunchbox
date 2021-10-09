package lunch

import (
	"flag"
	"fmt"
	"time"
)

func Date(args []string) error {
	var (
		dateCmd  = flag.NewFlagSet("date", flag.ExitOnError)
		timezone = dateCmd.String("location", "", "timezone location (e.g. UTC). Defaults to local time")
		layout   = dateCmd.String("layout", time.RFC3339, "layout according to the golang time package")
	)
	err := dateCmd.Parse(args)
	if err != nil {
		return err
	}
	n := time.Now()
	if *timezone != "" {
		loc, err := time.LoadLocation(*timezone)
		if err != nil {
			return err
		}
		n = n.In(loc)
	}
	fmt.Println(n.Format(*layout))
	return nil
}
