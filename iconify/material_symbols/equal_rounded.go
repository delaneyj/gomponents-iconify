package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EqualRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 17q-.625 0-1.063-.438T4 15.5q0-.625.438-1.063T5.5 14h13q.625 0 1.063.438T20 15.5q0 .625-.438 1.063T18.5 17h-13Zm0-7q-.625 0-1.063-.438T4 8.5q0-.625.438-1.063T5.5 7h13q.625 0 1.063.438T20 8.5q0 .625-.438 1.063T18.5 10h-13Z"/>`),
		g.Group(children),
	)
}