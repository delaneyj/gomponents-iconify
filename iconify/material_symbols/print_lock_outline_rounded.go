package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintLockOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 21q-.425 0-.713-.288T16 20v-3q0-.425.288-.713T17 16v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1q.425 0 .713.288T22 17v3q0 .425-.288.713T21 21h-4Zm1-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1ZM4 10h16H4Zm2 10v-3H3q-.425 0-.713-.288T2 16v-5q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.75q-.45-.25-.95-.425t-1.05-.25V11q0-.425-.288-.713T19 10H5q-.425 0-.713.288T4 11v4h2v-1q0-.425.288-.713T7 13h7.5q-.4.425-.725.925T13.25 15H8v4h5.1q.175.55.425 1.05t.6.95H7q-.425 0-.713-.288T6 20ZM6 8V5q0-.825.588-1.413T8 3h8q.825 0 1.413.588T18 5v3h-2V5H8v3H6Z"/>`),
		g.Group(children),
	)
}