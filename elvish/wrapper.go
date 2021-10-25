package elvish

import (
	"context"
	"embed"
	"io/fs"
)

//go:embed *
var plugins embed.FS

func List(_ context.Context) ([]string, error) {
	lst := []string{}
	if err := fs.WalkDir(plugins, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() != "wrapper.go" {
			lst = append(lst, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return lst, nil
}
