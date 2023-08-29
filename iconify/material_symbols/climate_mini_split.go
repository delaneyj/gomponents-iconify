package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClimateMiniSplit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 12q-.825 0-1.413-.588T2 10V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6q0 .825-.588 1.413T20 12h-2V8q0-.825-.588-1.413T16 6H8q-.825 0-1.413.588T6 8v4H4Zm3 2h2q0 2.075-1.463 3.538T4 19v-2q1.25 0 2.125-.875T7 14Zm1-2V8h8v4H8Zm3 8v-6h2v6h-2Zm4-6h2q0 1.25.875 2.125T20 17v2q-2.075 0-3.538-1.463T15 14Z"/>`),
		g.Group(children),
	)
}