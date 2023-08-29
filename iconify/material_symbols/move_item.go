package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoveItem(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.15 13H8v-2h12.15L18.6 9.45L20 8l4 4l-4 4l-1.4-1.45L20.15 13ZM15 9V5H5v14h10v-4h2v4q0 .825-.588 1.413T15 21H5q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h10q.825 0 1.413.588T17 5v4h-2Z"/>`),
		g.Group(children),
	)
}