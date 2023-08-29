package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottomDrawerOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v14q0 .825-.588 1.413T19 21H5Zm0-8.45q.45-.275.963-.413T7 12h10q.525 0 1.038.138t.962.412V5H5v7.55ZM5 19h14v-3q0-.825-.588-1.413T17 14H7q-.825 0-1.413.588T5 16v3Zm0 0h14H5Z"/>`),
		g.Group(children),
	)
}