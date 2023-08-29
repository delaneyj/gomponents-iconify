package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManageSearch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19v-2h10v2H2Zm0-5v-2h5v2H2Zm0-5V7h5v2H2Zm18.6 10l-3.85-3.85q-.6.425-1.313.638T14 16q-2.075 0-3.538-1.463T9 11q0-2.075 1.463-3.538T14 6q2.075 0 3.538 1.463T19 11q0 .725-.213 1.438t-.637 1.312L22 17.6L20.6 19ZM14 14q1.25 0 2.125-.875T17 11q0-1.25-.875-2.125T14 8q-1.25 0-2.125.875T11 11q0 1.25.875 2.125T14 14Z"/>`),
		g.Group(children),
	)
}