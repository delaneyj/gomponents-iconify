package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelfPositionOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19V5q0-.825.588-1.413T5 3h15q.825 0 1.413.588T22 5v14q0 .825-.588 1.413T20 21H5Zm0-5v3h15v-3H5Zm12-2h3V5h-3v9ZM5 14h3V5H5v9Zm5 0h5V5h-5v9Z"/>`),
		g.Group(children),
	)
}