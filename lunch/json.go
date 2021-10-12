package lunch

import (
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"text/template"
)

// deprecated. see JQ
func JSON(ctx Context, args []string) error {
	var (
		jsonCmd = flag.NewFlagSet("json", flag.ExitOnError)
		query   = jsonCmd.String("query", "{{.|json}}", "JSON query (go template syntax)")
	//	pretty  = jsonCmd.Bool("pretty", false, "pretty print")
	)
	err := jsonCmd.Parse(args)
	if err != nil {
		return err
	}
	data, err := unmarshalInput(os.Stdin, 25600)
	if err != nil {
		return err
	}
	tmpl, err := template.New("main").Funcs(funcMap()).Parse(*query)
	if err != nil {
		return err
	}
	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		return err
	}
	return nil
}

func funcMap() template.FuncMap {
	var funcMap = template.FuncMap{
		"json": toJSON,
	}
	return funcMap
}

func toJSON(x interface{}) (string, error) {
	switch xT := x.(type) {
	case map[string]interface{}:
		bts, err := json.Marshal(xT)
		return string(bts), err
	}
	bts, err := json.Marshal(x)
	return string(bts), err
}

func unmarshalInput(input io.ReadCloser, maxBufferSize int) (map[string]interface{}, error) {
	defer input.Close()
	y, err := ioutil.ReadAll(io.LimitReader(input, int64(maxBufferSize)))
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(y, &data)
	return data, err
}
