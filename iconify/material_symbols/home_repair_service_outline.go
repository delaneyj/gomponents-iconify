package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeRepairServiceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 14ZM2 20V10q0-.825.588-1.413T4 8h3V6q0-.825.588-1.413T9 4h6q.825 0 1.413.588T17 6v2h3q.825 0 1.413.588T22 10v10H2Zm6-5v1H6v-1H4v3h16v-3h-2v1h-2v-1H8Zm-4-5v3h2v-1h2v1h8v-1h2v1h2v-3H4Zm5-2h6V6H9v2Z"/>`),
		g.Group(children),
	)
}