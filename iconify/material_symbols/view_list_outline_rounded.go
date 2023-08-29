package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewListOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7q0-.825.588-1.413T5 5h14q.825 0 1.413.588T21 7v10q0 .825-.588 1.413T19 19H5q-.825 0-1.413-.588T3 17Zm2-8h2V7H5v2Zm4 0h10V7H9v2Zm0 4h10v-2H9v2Zm0 4h10v-2H9v2Zm-4 0h2v-2H5v2Zm0-4h2v-2H5v2Z"/>`),
		g.Group(children),
	)
}