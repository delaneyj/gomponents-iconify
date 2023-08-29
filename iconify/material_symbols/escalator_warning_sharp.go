package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EscalatorWarningSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 6q-.825 0-1.413-.588T4.5 4q0-.825.588-1.413T6.5 2q.825 0 1.413.588T8.5 4q0 .825-.588 1.413T6.5 6ZM17 11q-.625 0-1.063-.438T15.5 9.5q0-.625.438-1.063T17 8q.625 0 1.063.438T18.5 9.5q0 .625-.438 1.063T17 11ZM4.5 22v-7H3V7h6.15l4.15 7.175L14.775 12H20v5h-1v5h-4v-7.1l-.775 1.1h-2.2L9.5 11.6V22h-5Z"/>`),
		g.Group(children),
	)
}