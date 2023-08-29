package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RipplesOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 19h14v-8.55q-.45.275-.963.413T17 11q-1.65 0-2.825-1.175T13 7q0-.525.138-1.038T13.55 5H5v14Zm0 2q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5ZM5 5v14V5Z"/>`),
		g.Group(children),
	)
}