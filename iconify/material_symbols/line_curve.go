package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineCurve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 20q0-2.9-1.1-5.45t-3-4.45q-1.9-1.9-4.45-3T4 6V4q3.3 0 6.212 1.263T15.3 8.7q2.175 2.175 3.438 5.088T20 20h-2Z"/>`),
		g.Group(children),
	)
}