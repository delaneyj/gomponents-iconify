package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rhombus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2.59 12.5l8.91-8.91l8.91 8.91l-8.91 8.91l-8.91-8.91M11.5 5L4 12.5l7.5 7.5l7.5-7.5L11.5 5Z"/>`),
		g.Group(children),
	)
}