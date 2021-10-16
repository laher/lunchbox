package lunch

import (
	"flag"
	"fmt"
	"sort"

	"github.com/joho/godotenv"
)

func init() {
	Register("dotenv", resolveDotenv)
}

func resolveDotenv(ctx Context, args []string) error {
	var (
		flagset = flag.NewFlagSet("dotenv", flag.ExitOnError)
		//fileFlag = flagset.String("file", ".env", "file to parse")
	)
	err := flagset.Parse(args)
	if err != nil {
		return err
	}

	envMap, err := godotenv.Read(flagset.Args()[0])
	if err != nil {
		return err
	}
	// unquoted for now (useful for elvish). For quoted strings, introduce a boolean flag and use godotenv.Marshal()
	lines := make([]string, 0, len(envMap))
	for k, v := range envMap {
		lines = append(lines, fmt.Sprintf(`%s=%v`, k, v))
	}
	sort.Strings(lines)
	for _, l := range lines {
		//resolved
		fmt.Println(l)
	}
	return nil
}
