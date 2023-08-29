package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v11l-5 5Zm-1-2v-2q0-.825.588-1.413T17 15h2V5H5v14h10Zm-4-3h2v-6h3V8H8v2h3v6Zm4 3ZM5 19V5v14Z"/>`),
		g.Group(children),
	)
}