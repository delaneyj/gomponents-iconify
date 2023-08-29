package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Options(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5 3.5a1.5 1.5 0 1 1-3.001-.001A1.5 1.5 0 0 1 5 3.5z"/><path fill="currentColor" d="M3.5 0C1.6 0 0 1.6 0 3.5S1.6 7 3.5 7S7 5.4 7 3.5S5.4 0 3.5 0zm0 6C2.1 6 1 4.9 1 3.5S2.1 1 3.5 1S6 2.1 6 3.5S4.9 6 3.5 6zm0 2C1.6 8 0 9.6 0 11.5S1.6 15 3.5 15S7 13.4 7 11.5S5.4 8 3.5 8zm0 6C2.1 14 1 12.9 1 11.5S2.1 9 3.5 9S6 10.1 6 11.5S4.9 14 3.5 14zM8 2h8v3H8V2zm0 8h8v3H8v-3z"/>`),
		g.Group(children),
	)
}