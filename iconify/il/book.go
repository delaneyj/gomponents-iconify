package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Book(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M116 564q-20 0-33 13t-14 34q0 19 14 32t33 14h417q9 0 16 6t7 17v12q0 24-17 41t-41 17H69q-29 0-49-21T0 680V78q0-29 20-50T69 8h452q14 0 25 10t10 25v498q0 10-7 16t-16 7H116zm46-313q0 12 12 12h208q12 0 12-12v-69q0-5-3-9t-9-3H174q-5 0-8 3t-4 9v69z"/>`),
		g.Group(children),
	)
}