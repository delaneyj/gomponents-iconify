package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomPanelCloseOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 11.5l4-4H8l4 4ZM19 3q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14ZM5 16v3h14v-3H5Zm14-2V5H5v9h14ZM5 16v3v-3Z"/>`),
		g.Group(children),
	)
}