package simple_icons_go

import (
	_ "embed" // Embed Import for Package Files
	"errors"
)

//go:embed SI_VERSION
var version string

type SimpleIcon struct {
	release Release
	slugs   []Slug
	icons   Icons
}

func Load() SimpleIcon {
	release := LoadRelease(version)
	return SimpleIcon{
		release: release,
		slugs:   release.GetSlugs(),
		icons:   release.GetIcons(),
	}
}

func (si SimpleIcon) Get(slug string) (Icon, error) {
	for _, s := range si.slugs {
		if s.Slug == slug {
			icon := si.icons.getByName(s.Name)
			svg, err := si.release.GetSvg(s.Slug)
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
