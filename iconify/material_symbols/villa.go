package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Villa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-9h8q0-.825.588-1.413T19 10q.825 0 1.413.588T21 12v9h-5v-5h-2v5H9Zm-6 0V8l13-5v7H7v11H3Z"/>`),
		g.Group(children),
	)
}