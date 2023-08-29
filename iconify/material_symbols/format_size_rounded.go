package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatSizeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M14 7h-3.5q-.625 0-1.063-.438T9 5.5q0-.625.438-1.063T10.5 4h10q.625 0 1.063.438T22 5.5q0 .625-.438 1.063T20.5 7H17v11.5q0 .625-.438 1.063T15.5 20q-.625 0-1.063-.438T14 18.5V7Zm-9 5H3.5q-.625 0-1.063-.438T2 10.5q0-.625.438-1.063T3.5 9h6q.625 0 1.063.438T11 10.5q0 .625-.438 1.063T9.5 12H8v6.5q0 .625-.438 1.063T6.5 20q-.625 0-1.063-.438T5 18.5V12Z"/>`),
		g.Group(children),
	)
}