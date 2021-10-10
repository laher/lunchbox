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
		slurpFlag = jqCmd.Bool("slurp", false, "slurp")
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
		if *slurpFlag {

		} else {
			fmt.Printf("%#v\n", v)
		}
	}
	return nil
}
