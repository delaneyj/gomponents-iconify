package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toolbar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-9h18v9q0 .825-.588 1.413T19 21H5ZM3 8V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v3H3Z"/>`),
		g.Group(children),
	)
}