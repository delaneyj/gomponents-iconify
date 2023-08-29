package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.93 15.644L3.34 8.796a1.162 1.162 0 0 1-.002-1.641L7.904.323a1.162 1.162 0 0 1 1.642.002l4.54 6.85a1.16 1.16 0 0 1 .004 1.641l-4.518 6.83s-1.187.451-1.642-.002z"/>`),
		g.Group(children),
	)
}