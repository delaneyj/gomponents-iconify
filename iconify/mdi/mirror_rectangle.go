package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MirrorRectangle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.29 10.28l3.24-3.25l1.06 1.06l-3.24 3.25l-1.06-1.06m.41 4.33l5.66-5.66L15.42 10l-5.66 5.67l-1.06-1.06M18 3v18H6V3h12m2-2H4v22h16V1Z"/>`),
		g.Group(children),
	)
}