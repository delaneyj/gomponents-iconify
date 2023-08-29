package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EventUpcomingOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 22h-3q-.425 0-.713-.288T15 21q0-.425.288-.713T16 20h3V10H5v3q0 .425-.288.713T4 14q-.425 0-.713-.288T3 13V6q0-.825.588-1.413T5 4h1V3q0-.425.288-.713T7 2q.425 0 .713.288T8 3v1h8V3q0-.425.288-.713T17 2q.425 0 .713.288T18 3v1h1q.825 0 1.413.588T21 6v14q0 .825-.588 1.413T19 22Zm-9.825-2H2q-.425 0-.713-.288T1 19q0-.425.288-.713T2 18h7.175L7.3 16.1q-.275-.275-.287-.688T7.3 14.7q.275-.275.7-.275t.7.275l3.6 3.6q.3.3.3.7t-.3.7l-3.6 3.6q-.275.275-.687.288T7.3 23.3q-.275-.275-.275-.7t.275-.7L9.175 20ZM5 8h14V6H5v2Zm0 0V6v2Z"/>`),
		g.Group(children),
	)
}