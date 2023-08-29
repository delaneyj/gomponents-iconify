package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreakingNewsAltOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6 17h5v-2H6v2Zm10 0h2v-2h-2v2ZM6 13h5v-2H6v2Zm10 0h2V7h-2v6ZM6 9h5V7H6v2ZM4 21q-.825 0-1.413-.588T2 19V5q0-.825.588-1.413T4 3h16q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H4Z"/>`),
		g.Group(children),
	)
}