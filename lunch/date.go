package lunch

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/lestrrat-go/strftime"
)

func Date(args []string) error {
	var (
		dateCmd  = flag.NewFlagSet("date", flag.ExitOnError)
		timezone = dateCmd.String("location", "", "timezone location (e.g. UTC). Defaults to local time")
		layout   = dateCmd.String("layout", "", "layout according to the golang time package")
		format   = dateCmd.String("format", "%d/%m %H:%m:%S", "format according to strftime (%Y%M%D ...)")
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

	if *layout != "" {
		fmt.Println(n.Format(*layout))
	} else {
		f, err := strftime.New(*format)
		if err != nil {
			log.Println(err.Error())
		}
		buf := &bytes.Buffer{}
		if err := f.Format(buf, n); err != nil {
			log.Println(err.Error())
		}
		fmt.Println(buf.String())
	}
	return nil
}
