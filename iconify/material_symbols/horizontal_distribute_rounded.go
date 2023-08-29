package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HorizontalDistributeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21V3q0-.425.288-.713T3 2q.425 0 .713.288T4 3v18q0 .425-.288.713T3 22Zm9-5q-.625 0-1.063-.438T10.5 15.5v-7q0-.625.438-1.063T12 7q.625 0 1.063.438T13.5 8.5v7q0 .625-.438 1.063T12 17Zm9 5q-.425 0-.713-.288T20 21V3q0-.425.288-.713T21 2q.425 0 .713.288T22 3v18q0 .425-.288.713T21 22Z"/>`),
		g.Group(children),
	)
}