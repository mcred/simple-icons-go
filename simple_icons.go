package simple_icons_go

import (
	"embed"
	_ "embed" // Embed Import for Package Files
	"encoding/json"
	"errors"
	"strings"
)

//go:embed vendor/simple-icons/slugs.md
var rawSlugs string

//go:embed vendor/simple-icons/_data/simple-icons.json
var rawData []byte

//go:embed vendor/simple-icons/icons/*
var svgs embed.FS

type SimpleIcon struct{}

var icons Icons
var slugs []Slug
var iconDir string

func init() {
	for _, s := range strings.Split(rawSlugs, "\n") {
		if strings.HasPrefix(s, "|") {
			split := strings.Split(s, "|")
			if split[1] != "Brand name" && split[1] != ":---" {
				slug := Slug{Name: split[1][2 : len(split[1])-2], Slug: split[2][2 : len(split[2])-2]}
				slugs = append(slugs, slug)
			}
		}
	}
	err := json.Unmarshal(rawData, &icons)
	if err != nil {
		panic(err)
	}
}

func (si SimpleIcon) Get(slug string) (Icon, error) {
	for _, s := range slugs {
		if s.Slug == slug {
			icon := icons.getByName(s.Name)
			svg, err := svgs.ReadFile("vendor/simple-icons/icons/" + s.Slug + ".svg")
			if err != nil {
				goto end
			}
			icon.Slug = s.Slug
			icon.Svg = string(svg)
			return icon, nil
		}
	}
end:
	return Icon{}, errors.New("unable to load icon")
}
