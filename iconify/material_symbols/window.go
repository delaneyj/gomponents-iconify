package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Window(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 13h8v6q0 .825-.588 1.413T19 21h-6v-8Zm0-2V3h6q.825 0 1.413.588T21 5v6h-8Zm-2 0H3V5q0-.825.588-1.413T5 3h6v8Zm0 2v8H5q-.825 0-1.413-.588T3 19v-6h8Z"/>`),
		g.Group(children),
	)
}