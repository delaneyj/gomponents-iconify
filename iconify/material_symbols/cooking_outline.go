package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CookingOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-3H2v-2h7q.825 0 1.413.588T11 18v3H9Zm4 0v-3q0-.825.588-1.413T15 16h7v2h-7v3h-2Zm-7-6q-1.25 0-2.125-.875T3 12V8h18v4q0 1.25-.875 2.125T18 15H6Zm0-2h12q.425 0 .713-.288T19 12v-2H5v2q0 .425.288.713T6 13ZM3 7V5h6V4q0-.425.288-.713T10 3h4q.425 0 .713.288T15 4v1h6v2H3Zm2 6v-3v3Z"/>`),
		g.Group(children),
	)
}