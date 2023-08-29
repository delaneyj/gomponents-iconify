package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalDistributeRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22q-.425 0-.713-.288T2 21q0-.425.288-.713T3 20h18q.425 0 .713.288T22 21q0 .425-.288.713T21 22H3Zm5.5-8.5q-.625 0-1.063-.438T7 12q0-.625.438-1.063T8.5 10.5h7q.625 0 1.063.438T17 12q0 .625-.438 1.063T15.5 13.5h-7ZM3 4q-.425 0-.713-.288T2 3q0-.425.288-.713T3 2h18q.425 0 .713.288T22 3q0 .425-.288.713T21 4H3Z"/>`),
		g.Group(children),
	)
}