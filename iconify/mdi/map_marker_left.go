package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapMarkerLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 11.5a2.5 2.5 0 0 1 0-5a2.5 2.5 0 0 1 0 5M8 9c0 5.25 7 13 7 13s7-7.75 7-13c0-3.87-3.13-7-7-7S8 5.13 8 9M6 7l-5 5l5 5V7Z"/>`),
		g.Group(children),
	)
}