package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditAttributesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-2.075 0-3.538-1.463T2 12q0-2.075 1.463-3.538T7 7h10q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17H7Zm0-2h10q1.25 0 2.125-.875T20 12q0-1.25-.875-2.125T17 9H7q-1.25 0-2.125.875T4 12q0 1.25.875 2.125T7 15Zm1.05-.7l3.55-3.55l-1.05-1.05l-2.5 2.5l-1-1L6 12.25l2.05 2.05ZM12 12Z"/>`),
		g.Group(children),
	)
}