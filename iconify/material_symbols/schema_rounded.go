package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SchemaRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 23q-.625 0-1.063-.438T4 21.5v-3q0-.625.438-1.063T5.5 17h1v-2h-1q-.625 0-1.063-.438T4 13.5v-3q0-.625.438-1.063T5.5 9h1V7h-1q-.625 0-1.063-.438T4 5.5v-3q0-.625.438-1.063T5.5 1h4q.625 0 1.063.438T11 2.5v3q0 .625-.438 1.063T9.5 7h-1v2h1q.625 0 1.063.438T11 10.5v.5h3v-.5q0-.625.438-1.063T15.5 9h4q.625 0 1.063.438T21 10.5v3q0 .625-.438 1.063T19.5 15h-4q-.625 0-1.063-.438T14 13.5V13h-3v.5q0 .625-.438 1.063T9.5 15h-1v2h1q.625 0 1.063.438T11 18.5v3q0 .625-.438 1.063T9.5 23h-4Z"/>`),
		g.Group(children),
	)
}