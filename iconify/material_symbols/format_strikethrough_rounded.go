package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatStrikethroughRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 14q-.425 0-.713-.288T2 13q0-.425.288-.713T3 12h18q.425 0 .713.288T22 13q0 .425-.288.713T21 14H3Zm7.5-4V7h-4q-.625 0-1.063-.438T5 5.5q0-.625.438-1.063T6.5 4h11q.625 0 1.063.438T19 5.5q0 .625-.438 1.063T17.5 7h-4v3h-3Zm0 6h3v2.5q0 .625-.438 1.063T12 20q-.625 0-1.063-.438T10.5 18.5V16Z"/>`),
		g.Group(children),
	)
}