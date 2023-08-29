package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TeamDashboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.825 0-1.413-.588T3 19v-6.25h7V21H5Zm7 0v-8.25h9V19q0 .825-.588 1.413T19 21h-7ZM3 10.75V5q0-.825.588-1.413T5 3h14q.825 0 1.413.588T21 5v5.75H3Z"/>`),
		g.Group(children),
	)
}