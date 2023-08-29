package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRepairService(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 8h6V6H9v2ZM2 20v-5h4v1h2v-1h8v1h2v-1h4v5H2Zm0-6v-4q0-.825.588-1.413T4 8h3V6q0-.825.588-1.413T9 4h6q.825 0 1.413.588T17 6v2h3q.825 0 1.413.588T22 10v4h-4v-2h-2v2H8v-2H6v2H2Z"/>`),
		g.Group(children),
	)
}