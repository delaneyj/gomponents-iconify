package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TableRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19ZM5 9h14V5H5v4Zm5.325 5h3.35v-3h-3.35v3Zm0 5h3.35v-3h-3.35v3ZM5 14h3.325v-3H5v3Zm10.675 0H19v-3h-3.325v3ZM5 19h3.325v-3H5v3Zm10.675 0H19v-3h-3.325v3Z"/>`),
		g.Group(children),
	)
}