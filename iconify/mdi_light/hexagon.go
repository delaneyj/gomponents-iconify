package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.59 21l-4.9-8.5L6.6 4h9.81l4.91 8.5l-4.91 8.5H6.59m9.24-16H7.18l-4.34 7.5L7.17 20h8.66l4.33-7.5L15.83 5Z"/>`),
		g.Group(children),
	)
}