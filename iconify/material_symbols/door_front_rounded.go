package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorFrontRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 21q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v14h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Zm10-8q.425 0 .713-.288T15 12q0-.425-.288-.713T14 11q-.425 0-.713.288T13 12q0 .425.288.713T14 13Z"/>`),
		g.Group(children),
	)
}