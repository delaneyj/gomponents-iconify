package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewDayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17q-.825 0-1.413-.588T3 15V9q0-.825.588-1.413T5 7h14q.825 0 1.413.588T21 9v6q0 .825-.588 1.413T19 17H5Zm0-2h14V9H5v6ZM3 5V3h18v2H3Zm0 16v-2h18v2H3ZM5 9v6v-6Z"/>`),
		g.Group(children),
	)
}