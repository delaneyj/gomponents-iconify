package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UnknownMed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.5 13q-.625 0-1.063-.438T2 11.5q0-.625.438-1.063T3.5 10h6q.625 0 1.063.438T11 11.5q0 .625-.438 1.063T9.5 13h-6Zm11 0q-.625 0-1.063-.438T13 11.5q0-.625.438-1.063T14.5 10h6q.625 0 1.063.438T22 11.5q0 .625-.438 1.063T20.5 13h-6Z"/>`),
		g.Group(children),
	)
}