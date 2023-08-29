package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 11.5a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5M9 2C5.13 2 2 5.13 2 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7m9 15l5-5l-5-5v10Z"/>`),
		g.Group(children),
	)
}