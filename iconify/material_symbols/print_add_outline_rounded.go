package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintAddOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10h16H4Zm2 10v-3H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.45-.25-.95-.425t-1.05-.25V11q0-.425-.288-.713T19 10H5q-.425 0-.713.288T4 11v4h2v-1q0-.425.288-.713T7 13h7.5q-.4.425-.725.925T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H7q-.425 0-.713-.288T6 20ZM6 8V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v3h-2V5H8v3H6Zm12 10h-2q-.425 0-.713-.288T15 17q0-.425.288-.713T16 16h2v-2q0-.425.288-.713T19 13q.425 0 .713.288T20 14v2h2q.425 0 .713.288T23 17q0 .425-.288.713T22 18h-2v2q0 .425-.288.713T19 21q-.425 0-.713-.288T18 20v-2Z"/>`),
		g.Group(children),
	)
}