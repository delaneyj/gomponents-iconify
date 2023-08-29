package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignHorizontalRightRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 22q-.425 0-.713-.288T20 21V3q0-.425.288-.713T21 2q.425 0 .713.288T22 3v18q0 .425-.288.713T21 22ZM9.5 17q-.625 0-1.063-.438T8 15.5q0-.625.438-1.063T9.5 14h7q.625 0 1.063.438T18 15.5q0 .625-.438 1.063T16.5 17h-7Zm-6-7q-.625 0-1.063-.438T2 8.5q0-.625.438-1.063T3.5 7h13q.625 0 1.063.438T18 8.5q0 .625-.438 1.063T16.5 10h-13Z"/>`),
		g.Group(children),
	)
}