package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TitleRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.5 7h-4q-.625 0-1.063-.438T5 5.5q0-.625.438-1.063T6.5 4h11q.625 0 1.063.438T19 5.5q0 .625-.438 1.063T17.5 7h-4v11.5q0 .625-.438 1.063T12 20q-.625 0-1.063-.438T10.5 18.5V7Z"/>`),
		g.Group(children),
	)
}