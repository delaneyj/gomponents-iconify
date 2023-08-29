package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterMedium(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.125 12.15q1.175-.575 2.438-.863T11.124 11q.75 0 1.488.1t1.462.3q1.25.35 1.913.475T17.4 12h.475l.875-8H5.25l.875 8.15Zm.85 9.85q-.775 0-1.338-.5T5 20.225L3 2h18l-2 18.225q-.075.775-.638 1.275t-1.337.5H6.975Z"/>`),
		g.Group(children),
	)
}