package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTravelOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 17v2h16v-2H4ZM7 6V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v2h3q.825 0 1.413.588T22 8v11q0 .825-.588 1.413T20 21H4q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6h3Zm-3 8h16V8h-3v2h-2V8H9v2H7V8H4v6Zm5-8h6V4H9v2ZM4 19V8v2v-2v2v-2v11Z"/>`),
		g.Group(children),
	)
}