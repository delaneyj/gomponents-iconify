package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m19.25 16.45l-1.5-1.55q1-.275 1.625-1.063T20 12q0-1.25-.875-2.125T17 9h-4V7h4q2.075 0 3.538 1.463T22 12q0 1.425-.738 2.625T19.25 16.45ZM15.85 13l-2-2H16v2h-.15Zm3.95 9.6L1.4 4.2l1.4-1.4l18.4 18.4l-1.4 1.4ZM11 17H7q-2.075 0-3.538-1.463T2 12q0-1.725 1.05-3.075t2.7-1.775L7.6 9H7q-1.25 0-2.125.875T4 12q0 1.25.875 2.125T7 15h4v2Zm-3-4v-2h1.625l1.975 2H8Z"/>`),
		g.Group(children),
	)
}