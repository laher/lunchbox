package lunch

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/itchyny/gojq"
)

func JQ(args []string) error {

	var (
		jqCmd     = flag.NewFlagSet("jq", flag.ExitOnError)
		queryFlag = jqCmd.String("query", ".", "jq-style query (gojq variant)")
	)
	query, err := gojq.Parse(*queryFlag)
	if err != nil {
		return err
	}
	input := make(map[string]interface{})
	if err := json.NewDecoder(os.Stdin).Decode(&input); err != nil {
		return err
	}
	iter := query.Run(input) // or query.RunWithContext
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return err
		}
		fmt.Printf("%#v\n", v)
	}
	return nil
}
