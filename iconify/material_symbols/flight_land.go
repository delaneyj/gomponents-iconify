package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightLand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.85 15.9L3 11.5V6l1.5.5l.7 2.1L10 9.95V2l2 .5l2.75 8.75l5 1.4q.575.15.913.613T21 14.3q0 .825-.675 1.325t-1.475.275ZM3 21v-2h18v2H3Z"/>`),
		g.Group(children),
	)
}