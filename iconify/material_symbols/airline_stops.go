package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineStops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 19v-2h2q-.375-3.45-2.95-5.725T2 9V7q3.225 0 5.925 1.7T12 13.3q.95-2.025 2.5-3.588T17.975 7H14V5h7v7h-2V8.7q-2.325 1.425-4 3.525T13 17h2v2H9Z"/>`),
		g.Group(children),
	)
}