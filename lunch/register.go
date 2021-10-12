package lunch

import (
	"context"
	"os"
)

type Context struct {
	Ctx context.Context
	In  *os.File
	Out *os.File
	Err *os.File
}

type Lunch func(ctx Context, args []string) error

var lunches = map[string]Lunch{}

func Register(name string, meal Lunch) {
	lunches[name] = meal
}

func GetNames() []string {
	names := make([]string, 0, len(lunches))
	for k := range lunches {
		names = append(names, k)
	}
	return names
}

func Get(name string) (Lunch, bool) {
	l, ok := lunches[name]
	return l, ok
}
