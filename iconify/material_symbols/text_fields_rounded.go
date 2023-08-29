package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextFieldsRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.5 20q-.625 0-1.063-.438T7 18.5V7H3.5q-.625 0-1.063-.438T2 5.5q0-.625.438-1.063T3.5 4h10q.625 0 1.063.438T15 5.5q0 .625-.438 1.063T13.5 7H10v11.5q0 .625-.438 1.063T8.5 20Zm9 0q-.625 0-1.063-.438T16 18.5V12h-1.5q-.625 0-1.063-.438T13 10.5q0-.625.438-1.063T14.5 9h6q.625 0 1.063.438T22 10.5q0 .625-.438 1.063T20.5 12H19v6.5q0 .625-.438 1.063T17.5 20Z"/>`),
		g.Group(children),
	)
}