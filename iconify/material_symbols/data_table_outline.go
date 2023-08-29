package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataTableOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5ZM5 8.325h14V5H5v3.325Zm0 5.35h14v-3.35H5v3.35ZM5 19h14v-3.325H5V19ZM6 7.65v-2h2v2H6ZM6 13v-2h2v2H6Zm0 5.35v-2h2v2H6Z"/>`),
		g.Group(children),
	)
}