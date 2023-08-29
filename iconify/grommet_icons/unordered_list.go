package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnorderedList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.94 6.42H24v1.75H5.94zm0 5.29H24v1.75H5.94zm0 5.28H24v1.75H5.94z"/><circle cx="1.85" cy="7.29" r="1.52" fill="currentColor"/><circle cx="1.85" cy="12.58" r="1.52" fill="currentColor"/><circle cx="1.85" cy="17.87" r="1.52" fill="currentColor"/>`),
		g.Group(children),
	)
}