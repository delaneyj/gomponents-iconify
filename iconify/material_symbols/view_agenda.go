package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewAgenda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 11q-.825 0-1.413-.588T3 9V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v4q0 .825-.588 1.413T19 11H5Zm0 10q-.825 0-1.413-.588T3 19v-4q0-.825.588-1.413T5 13h14q.825 0 1.413.588T21 15v4q0 .825-.588 1.413T19 21H5Z"/>`),
		g.Group(children),
	)
}