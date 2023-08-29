package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewWeekOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm9-2h2.5V6H13v12Zm-4.5 0H11V6H8.5v12ZM4 18h2.5V6H4v12Zm13.5 0H20V6h-2.5v12Z"/>`),
		g.Group(children),
	)
}