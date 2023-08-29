package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-2h18v2H3ZM3 6V4h18v2H3Zm2 10q-.825 0-1.413-.588T3 14v-4q0-.825.588-1.413T5 8h14q.825 0 1.413.588T21 10v4q0 .825-.588 1.413T19 16H5Z"/>`),
		g.Group(children),
	)
}