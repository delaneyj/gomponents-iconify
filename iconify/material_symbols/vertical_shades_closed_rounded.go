package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalShadesClosedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21q-.425 0-.713-.288T2 20q0-.425.288-.713T3 19h1V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14h1q.425 0 .713.288T22 20q0 .425-.288.713T21 21H3Zm3-2h1.5V5H6v14Zm3.5 0H11V5H9.5v14Zm3.5 0h1.5V5H13v14Zm3.5 0H18V5h-1.5v14Z"/>`),
		g.Group(children),
	)
}