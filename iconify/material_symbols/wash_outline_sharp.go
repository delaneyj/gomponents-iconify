package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17.5 8q-.975 0-1.738-.763T15 5.5q0-.875.575-1.938T17.5 1q1.35 1.5 1.925 2.563T20 5.5q0 .975-.763 1.738T17.5 8ZM19 22H2V10.975l9.625-6.25L13.25 6.35L11.3 9.5H20v2H7.7l1.925-3.1L4 12.05V20h15v2Zm-7-7v-2h10v2H12Zm0 3.5v-2h9v2h-9Zm-4-3.05Z"/>`),
		g.Group(children),
	)
}