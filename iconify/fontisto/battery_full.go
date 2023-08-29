package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryFull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.12 17.6A3.12 3.12 0 0 1 0 14.48V8.72A3.12 3.12 0 0 1 3.12 5.6h16.011a3.12 3.12 0 0 1 3.12 3.12v5.76a3.12 3.12 0 0 1-3.12 3.12zM1.108 8.72v5.76a2.012 2.012 0 0 0 2.011 2.01h16.012a2.011 2.011 0 0 0 2.009-2.01V8.72a2.012 2.012 0 0 0-2.009-2.01H3.119c-1.11 0-2.01.9-2.011 2.01zm2.679 6.4a1.312 1.312 0 0 1-1.307-1.308V9.385a1.31 1.31 0 0 1 1.305-1.31h14.678a1.31 1.31 0 0 1 1.311 1.31v4.427a1.316 1.316 0 0 1-1.311 1.308zm18.87-5.41a2 2 0 0 1 .014 3.773l-.014.004z"/>`),
		g.Group(children),
	)
}