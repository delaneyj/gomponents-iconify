package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableChartOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h15q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21ZM5 8h15V5H5v3Zm3 2H5v9h3v-9Zm9 0v9h3v-9h-3Zm-2 0h-5v9h5v-9Z"/>`),
		g.Group(children),
	)
}