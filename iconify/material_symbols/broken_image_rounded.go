package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrokenImageRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-6.6l2.3 2.3q.3.3.7.3t.7-.3l3.3-3.3l3.3 3.3q.3.3.7.3t.7-.3l3.3-3.3l3 3V19q0 .825-.588 1.413T19 21H5ZM5 3h14q.825 0 1.413.588T21 5v6.575l-2.3-2.3q-.3-.3-.7-.3t-.7.3l-3.3 3.3l-3.3-3.3q-.3-.3-.7-.3t-.7.3l-3.3 3.3l-3-3V5q0-.825.588-1.413T5 3Z"/>`),
		g.Group(children),
	)
}