package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatLineSpacing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 6v1H10V6h11m0 6v1H10v-1h11m0 6v1H10v-1h11M5 19.25L7.25 17l.75.66l-3.5 3.5l-3.5-3.5l.75-.66L4 19.25V5.75L1.75 8L1 7.34l3.5-3.5L8 7.34L7.25 8L5 5.75v13.5Z"/>`),
		g.Group(children),
	)
}