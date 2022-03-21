package simple_icons_go

import (
	_ "embed" // Embed Import for Package Files
)

//go:embed SI_VERSION
var version string

type SimpleIcon struct {
	release Release
	slugs   []Slug
	icons   Icons
}

func Load() (SimpleIcon, error) {
	release, err := LoadRelease(version)
	slugs, err := release.GetSlugs()
	icons, err := release.GetIcons()
	return SimpleIcon{
		release: release,
		slugs:   slugs,
		icons:   icons,
	}, err
}

func (si SimpleIcon) Get(slug string) (Icon, error) {
	icon := Icon{}
	for _, s := range si.slugs {
		if s.Slug == slug {
			icon = si.icons.getByName(s.Name)
			svg, err := si.release.GetSvg(s.Slug)
			if err != nil {
				return icon, err
			}
			icon.Slug = s.Slug
			icon.Svg = string(svg)
		}
	}
	return icon, nil
}
