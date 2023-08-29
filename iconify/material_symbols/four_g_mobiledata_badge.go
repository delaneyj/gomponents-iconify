package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourGMobiledataBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.825 0-1.413-.588T1 19V5q0-.825.588-1.413T3 3h18q.825 0 1.413.588T23 5v14q0 .825-.588 1.413T21 21H3Zm17-10h-3.5v2H18v2h-3V9h5q0-.825-.588-1.413T18 7h-3q-.825 0-1.413.588T13 9v6q0 .825.588 1.413T15 17h3q.825 0 1.413-.588T20 15v-4ZM8 17h2v-3h2v-2h-2V7H8v5H6V7H4v7h4v3Z"/>`),
		g.Group(children),
	)
}