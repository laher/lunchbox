package xbarplugins

import (
	"context"
	"embed"
	"io/fs"
)

//go:embed *
var plugins embed.FS

func List(ctx context.Context) ([]string, error) {
	pluginList := []string{}
	if err := fs.WalkDir(plugins, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && d.Name() != "wrapper.go" {
			pluginList = append(pluginList, path)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return pluginList, nil
}

func Open(ctx context.Context, filename string) (fs.File, error) {
	return plugins.Open(filename)
}
