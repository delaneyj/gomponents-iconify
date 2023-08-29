package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatHTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 17V7h2v4h4V7h2v10H9v-4H5v4H3Zm10 0v-4q0-.825.588-1.413T15 11h4V9h-6V7h6q.825 0 1.413.588T21 9v2q0 .825-.588 1.413T19 13h-4v2h6v2h-8Z"/>`),
		g.Group(children),
	)
}