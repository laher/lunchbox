package valinor

import (
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"src.elv.sh/pkg/buildinfo"
	"src.elv.sh/pkg/daemon/client"
	"src.elv.sh/pkg/eval"
	"src.elv.sh/pkg/parse"
	"src.elv.sh/pkg/prog"
	"src.elv.sh/pkg/shell"
)

func ElvishPrompt(args []string) error {
	// prepend an empty string
	return elvishPrompt(append([]string{""}, args...))
}

func elvishPrompt(args []string) error {
	os.Exit(prog.Run(
		[3]*os.File{os.Stdin, os.Stdout, os.Stderr}, args,
		buildinfo.Program, daemonStub{}, shell.Program{ActivateDaemon: client.Activate}))
	return nil
}

func ElvishRunScript(args []string) error {
	var scriptCmd = flag.NewFlagSet("script", flag.ExitOnError)

	scriptCmd.Parse(os.Args[2:])
	if len(os.Args) > 2 {
		script := os.Args[2]
		godotenv.Load(filepath.Base(script) + ".env")
		bin := script
		_, err := elvishRunScript(bin, os.Stdout, os.Stderr, append([]string{""}, flag.Args()...))
		return err
	}
	return errors.New("required 'script' flag")
}

func elvishRunScript(bin string, out, stderr *os.File, args []string) ([]string, error) {
	f, err := os.ReadFile(bin)
	if err != nil {
		return []string{}, err
	}
	s := parse.Source{Name: bin, Code: string(f), IsFile: true}
	e := eval.NewEvaler()
	capture, fetcher, err := eval.StringCapturePort()
	if err != nil {
		return []string{}, err
	}
	cfg := eval.EvalCfg{
		PutInFg: true,
		Ports:   []*eval.Port{eval.DummyInputPort, capture, eval.DummyOutputPort}, // TODO stop using dummy output
	}

	/* load env
	variable := eval.MakeVarFromName(name)
	err := variable.Set(val)
	if err != nil {
		return err
	}
	*/

	err = e.Eval(s, cfg)
	if err != nil {
		return []string{}, err
	}

	return fetcher(), nil

	/*
		os.Exit(prog.Run(
			[3]*os.File{in, out, stderr}, args,
			buildinfo.Program, daemonStub{}, shell.Program{}))
	*/
	// ? prog.Composite(buildinfo.Program, daemonStub{}, shell.Program{})))
}

var errNoDaemon = errors.New("daemon is not supported in this build")

type daemonStub struct{}

func (daemonStub) ShouldRun(f *prog.Flags) bool { return f.Daemon }

func (daemonStub) Run(fds [3]*os.File, f *prog.Flags, args []string) error {
	return errNoDaemon
}