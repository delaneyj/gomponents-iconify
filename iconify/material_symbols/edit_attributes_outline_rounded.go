package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditAttributesOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 17q-2.075 0-3.538-1.463T2 12q0-2.075 1.463-3.538T7 7h10q2.075 0 3.538 1.463T22 12q0 2.075-1.463 3.538T17 17H7Zm0-2h10q1.25 0 2.125-.875T20 12q0-1.25-.875-2.125T17 9H7q-1.25 0-2.125.875T4 12q0 1.25.875 2.125T7 15Zm1.75-1.4l2.325-2.325q.225-.225.225-.525t-.225-.525Q10.85 10 10.55 10t-.525.225L8.05 12.2l-.475-.475Q7.35 11.5 7.05 11.5t-.525.225q-.225.225-.225.525t.225.525l.825.825q.3.3.7.3t.7-.3ZM12 12Z"/>`),
		g.Group(children),
	)
}