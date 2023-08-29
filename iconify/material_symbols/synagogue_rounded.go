package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SynagogueRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 8V7q0-.825.588-1.413T21 5q.825 0 1.413.588T23 7v1h-4ZM1 8V7q0-.825.588-1.413T3 5q.825 0 1.413.588T5 7v1H1Zm0 11V9h4v12H3q-.825 0-1.413-.588T1 19Zm5 2V8l4.725-3.925q.275-.225.6-.35T12 3.6q.35 0 .675.125t.6.35L18 8v13h-3q-.425 0-.713-.288T14 20v-4q0-.825-.588-1.413T12 14q-.825 0-1.413.588T10 16v4q0 .425-.288.713T9 21H6Zm13 0V9h4v10q0 .825-.588 1.413T21 21h-2Zm-7-9.5q.625 0 1.063-.438T13.5 10q0-.625-.438-1.063T12 8.5q-.625 0-1.063.438T10.5 10q0 .625.438 1.063T12 11.5Z"/>`),
		g.Group(children),
	)
}