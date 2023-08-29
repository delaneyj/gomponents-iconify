package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AutoAwesomeMosaic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h6v18Zm2-10V3h6q.825 0 1.413.588T21 5v6h-8Zm0 10v-8h8v6q0 .825-.588 1.413T19 21h-6Z"/>`),
		g.Group(children),
	)
}