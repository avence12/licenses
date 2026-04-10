package assets

import (
	"embed"
)

//go:embed *.txt
var fs embed.FS

type Asset struct {
	Name    string
	Content string
}

var Assets []Asset

func init() {
	entries, err := fs.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		content, err := fs.ReadFile(entry.Name())
		if err != nil {
			panic(err)
		}
		Assets = append(Assets, Asset{
			Name:    entry.Name(),
			Content: string(content),
		})
	}
}
