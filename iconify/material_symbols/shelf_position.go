package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelfPosition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 16v3q0 .825.588 1.413T5 21h15q.825 0 1.413-.588T22 19v-3H3Zm14-2h5V5q0-.825-.588-1.413T20 3h-3v11ZM3 14h5V3H5q-.825 0-1.413.588T3 5v9Zm7 0h5V3h-5v11Z"/>`),
		g.Group(children),
	)
}