package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14V7h9q1.65 0 2.825 1.175T22 11v3H9Zm-7 3v-2h20v2H2Zm3-3q-1.25 0-2.125-.875T2 11q0-1.25.875-2.125T5 8q1.25 0 2.125.875T8 11q0 1.25-.875 2.125T5 14Z"/>`),
		g.Group(children),
	)
}