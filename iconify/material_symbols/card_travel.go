package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardTravel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 6V4q0-.825.588-1.413T9 2h6q.825 0 1.413.588T17 4v2h3q.825 0 1.413.588T22 8v11q0 .825-.588 1.413T20 21H4q-.825 0-1.413-.588T2 19V8q0-.825.588-1.413T4 6h3Zm2 0h6V4H9v2ZM4 17h16v-3H4v3Z"/>`),
		g.Group(children),
	)
}