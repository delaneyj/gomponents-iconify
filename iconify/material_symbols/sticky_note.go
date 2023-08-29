package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v11l-5 5H5Zm6-5h2v-6h3V8H8v2h3v6Zm4 3l4-4h-2q-.825 0-1.413.588T15 17v2Z"/>`),
		g.Group(children),
	)
}