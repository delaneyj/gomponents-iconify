package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalendarViewMonthOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6v12q0 .825-.588 1.413T20 20H4Zm0-9h4V6H4v5Zm6 0h4V6h-4v5Zm6 0h4V6h-4v5Zm-8 7v-5H4v5h4Zm2 0h4v-5h-4v5Zm6 0h4v-5h-4v5Z"/>`),
		g.Group(children),
	)
}