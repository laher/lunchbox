package lunch

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/nojima/httpie-go"
	"github.com/nojima/httpie-go/flags"
	"github.com/nojima/httpie-go/input"
	"github.com/pkg/errors"
)

func init() {
	Register("http", HTTP)
}

func HTTP(ctx Context, args []string) error {

	// Parse flags
	args, usage, optionSet, err := flags.Parse(args)
	if err != nil {
		return err
	}
	inputOptions := optionSet.InputOptions
	exchangeOptions := optionSet.ExchangeOptions
	outputOptions := optionSet.OutputOptions

	// Parse positional arguments
	in, err := input.ParseArgs(args, os.Stdin, &inputOptions)
	if _, ok := errors.Cause(err).(*input.UsageError); ok {
		usage.PrintUsage(os.Stderr)
		return err
	}
	if err != nil {
		return err
	}

	// Send request and receive response
	status, err := httpie.Exchange(in, &exchangeOptions, &outputOptions)
	if err != nil {
		return err
	}

	if exchangeOptions.CheckStatus {
		os.Exit(getExitStatus(status))
	}

	return nil
}

func getExitStatus(statusCode int) int {
	if 300 <= statusCode && statusCode < 600 {
		return statusCode / 100
	}
	return 0
}

func HTTPx(args []string) error {
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
