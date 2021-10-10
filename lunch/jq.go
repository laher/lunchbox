package lunch

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/itchyny/gojq"
)

func JQ(ctx context.Context, args []string) error {

	var (
		jqCmd     = flag.NewFlagSet("jq", flag.ExitOnError)
		queryFlag = jqCmd.String("query", ".", "jq-style query (gojq variant)")
		rawFlag   = jqCmd.Bool("raw", false, "raw result (no quotes on a string)")
	)
	err := jqCmd.Parse(args)
	if err != nil {
		return err
	}
	query, err := gojq.Parse(*queryFlag)
	if err != nil {
		return err
	}
	input := make(map[string]interface{})
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		return err
	}
	iter := query.RunWithContext(ctx, input) // or query.RunWithContext
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}
		if *rawFlag {
			f := newEncoder(true, 2)
			m := &rawMarshaler{f}
			err := m.marshal(v, os.Stdout)
			if err != nil {
				return err
			}
			//fmt.Printf("%#v\n", string(b))
		} else {
			fmt.Printf("%#v\n", v)
		}
	}
	return nil
}
