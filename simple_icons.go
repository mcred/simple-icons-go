package simple_icons_go

import (
	_ "embed" // Embed Import for Package Files
	"errors"
)

//go:embed SI_VERSION
var version string

var release Release
var slugs []Slug
var icons Icons

type SimpleIcon struct{}

func init() {
	release = LoadRelease(version)
	slugs = release.GetSlugs()
	icons = release.GetIcons()
}

func (si SimpleIcon) Get(slug string) (Icon, error) {
	for _, s := range slugs {
		if s.Slug == slug {
			icon := icons.getByName(s.Name)
			svg, err := release.GetSvg(s.Slug)
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
