package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScatterChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21h17v-2H5V3H3v17a1 1 0 0 0 1 1z"/><circle cx="10" cy="8" r="2" fill="currentColor"/><circle cx="18" cy="12" r="2" fill="currentColor"/><circle cx="11.5" cy="13.5" r="1.5" fill="currentColor"/><circle cx="16.5" cy="6.5" r="1.5" fill="currentColor"/>`),
		g.Group(children),
	)
}