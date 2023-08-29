package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PrintLockOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 10h16H4Zm2 11v-4H2v-6q0-1.275.875-2.138T5 8h14q1.275 0 2.138.863T22 11v.8q-.45-.25-.95-.438T20 11.1q0-.425-.288-.762T19 10H5q-.425 0-.713.288T4 11v4h2v-2h8.55q-.4.425-.7.925T13.35 15H8v4h5.35q.175.55.513 1.05t.687.95H6ZM16 8V5H8v3H6V3h12v5h-2Zm1 13q-.425 0-.713-.288T16 20v-3q0-.425.288-.713T17 16v-1q0-.825.588-1.413T19 13q.825 0 1.413.588T21 15v1q.425 0 .713.288T22 17v3q0 .425-.288.713T21 21h-4Zm1-5h2v-1q0-.425-.288-.713T19 14q-.425 0-.713.288T18 15v1Z"/>`),
		g.Group(children),
	)
}