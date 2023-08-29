package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Synagogue(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8V7q0-.825.588-1.413T21 5q.825 0 1.413.588T23 7v1h-4ZM1 8V7q0-.825.588-1.413T3 5q.825 0 1.413.588T5 7v1H1Zm0 13V9h4v12H1Zm5 0V8l6-5l6 5v13h-4v-5q0-.825-.588-1.413T12 14q-.825 0-1.413.588T10 16v5H6Zm13 0V9h4v12h-4Zm-7-9.5q.625 0 1.063-.438T13.5 10q0-.625-.438-1.063T12 8.5q-.625 0-1.063.438T10.5 10q0 .625.438 1.063T12 11.5Z"/>`),
		g.Group(children),
	)
}