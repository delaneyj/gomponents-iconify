package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BorderAllOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19Zm10-6v6h6v-6h-6Zm0-2h6V5h-6v6Zm-2 0V5H5v6h6Zm0 2H5v6h6v-6Z"/>`),
		g.Group(children),
	)
}