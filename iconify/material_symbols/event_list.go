package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventList(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21q-.825 0-1.413-.588T14 19v-4q0-.825.588-1.413T16 13h4q.825 0 1.413.588T22 15v4q0 .825-.588 1.413T20 21h-4ZM2 18v-2h9v2H2Zm14-7q-.825 0-1.413-.588T14 9V5q0-.825.588-1.413T16 3h4q.825 0 1.413.588T22 5v4q0 .825-.588 1.413T20 11h-4ZM2 8V6h9v2H2Z"/>`),
		g.Group(children),
	)
}