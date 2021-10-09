package lunch

import (
	"flag"
	"fmt"
	"io"
	"net/http"
)

func HTTP(args []string) error {
	var (
		httpCmd              = flag.NewFlagSet("http", flag.ExitOnError)
		method               = httpCmd.String("method", "GET", "method of http call")
		printResponseHeaders = httpCmd.Bool("i", false, "print response headers")
	)
	err := httpCmd.Parse(args)
	if err != nil {
		return err
	}
	url := httpCmd.Args()[0]
	req, err := http.NewRequest(*method, url, nil)
	if err != nil {
		return err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(resp.Status)
	if *printResponseHeaders {
		for k, v := range resp.Header {
			for _, s := range v {
				fmt.Printf("%s: %s\n", k, s)
			}
		}
	}
	fmt.Println()
	fmt.Println(string(body))
	return nil
}
