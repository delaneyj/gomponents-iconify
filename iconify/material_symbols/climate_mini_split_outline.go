package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClimateMiniSplitOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 12H4q-.825 0-1.413-.588T2 10V4q0-.825.588-1.413T4 2h16q.825 0 1.413.588T22 4v6q0 .825-.588 1.413T20 12ZM4 19v-2q1.25 0 2.125-.875T7 14h2q0 2.075-1.463 3.538T4 19Zm16 0q-2.075 0-3.538-1.463T15 14h2q0 1.25.875 2.125T20 17v2Zm-9 1v-6h2v6h-2Zm9-10H4h16ZM6 10V8q0-.825.588-1.413T8 6h8q.825 0 1.413.588T18 8v2h-2V8H8v2H6Zm-2 0h16V4H4v6Z"/>`),
		g.Group(children),
	)
}