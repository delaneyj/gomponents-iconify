package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SportsCricketSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.5 17.7L1.625 7.825l4.2-4.2L15.7 13.5l-4.2 4.2Zm7.1 4.3l-4.25-4.25l1.4-1.4L20 20.6L18.6 22Zm-.1-13q-1.45 0-2.475-1.025T15 5.5q0-1.45 1.025-2.475T18.5 2q1.45 0 2.475 1.025T22 5.5q0 1.45-1.025 2.475T18.5 9Z"/>`),
		g.Group(children),
	)
}