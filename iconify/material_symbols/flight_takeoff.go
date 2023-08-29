package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlightTakeoff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21v-2h18v2H3Zm1.85-5L2 11.15l1.55-.3L5.3 12.4l4.8-1.3l-4.05-6.85L8 3.65l6.85 6.15l5-1.35q.8-.225 1.45.3t.65 1.4q0 .55-.338.975t-.862.575L4.85 16Z"/>`),
		g.Group(children),
	)
}