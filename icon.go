package simple_icons_go

type Icon struct {
	Title      string  `json:"title"`
	Slug       string  `json:"slug"`
	Hex        string  `json:"hex"`
	Source     string  `json:"source"`
	Svg        string  `json:"svg"`
	Guidelines string  `json:"guidelines"`
	License    License `json:"license"`
}

type Icons struct {
	Icons []Icon `json:"icons"`
}

func (i Icons) getByName(n string) Icon {
	for _, k := range i.Icons {
		if k.Title == n {
			return k
		}
	}
	return Icon{}
}
