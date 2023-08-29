package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalLeftRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21V3q0-.425.288-.713T3 2q.425 0 .713.288T4 3v18q0 .425-.288.713T3 22Zm4.5-5q-.625 0-1.063-.438T6 15.5q0-.625.438-1.063T7.5 14h7q.625 0 1.063.438T16 15.5q0 .625-.438 1.063T14.5 17h-7Zm0-7q-.625 0-1.063-.438T6 8.5q0-.625.438-1.063T7.5 7h13q.625 0 1.063.438T22 8.5q0 .625-.438 1.063T20.5 10h-13Z"/>`),
		g.Group(children),
	)
}