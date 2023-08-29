package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ApprovalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-6q0-.825.588-1.413T6 14h12q.825 0 1.413.588T20 16v6H4Zm2-4h12v-2H6v2Zm6-4L7 7q0-2.075 1.463-3.538T12 2q2.075 0 3.538 1.463T17 7l-5 7Zm0-2.8L15 7q0-1.25-.875-2.125T12 4q-1.25 0-2.125.875T9 7l3 4.2Zm0-3.6Z"/>`),
		g.Group(children),
	)
}