package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipExit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.075 16.5l1.425-1.425L15.4 12H18v-2h-6v6h2v-2.575l3.075 3.075ZM2 9V4h7v5H2Zm0 2h9V4h9q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4q-.825 0-1.413-.588T2 18v-7Z"/>`),
		g.Group(children),
	)
}