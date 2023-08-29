package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataObjectRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 20q-.425 0-.713-.288T14 19q0-.425.288-.713T15 18h2q.425 0 .713-.288T18 17v-2q0-.95.55-1.725t1.45-1.1v-.35q-.9-.325-1.45-1.1T18 9V7q0-.425-.288-.712T17 6h-2q-.425 0-.713-.288T14 5q0-.425.288-.713T15 4h2q1.25 0 2.125.875T20 7v2q0 .425.288.713T21 10t.713.288Q22 10.575 22 11v2q0 .425-.288.713T21 14t-.713.288Q20 14.575 20 15v2q0 1.25-.875 2.125T17 20h-2Zm-8 0q-1.25 0-2.125-.875T4 17v-2q0-.425-.288-.713T3 14t-.713-.288Q2 13.425 2 13v-2q0-.425.288-.713T3 10t.713-.288Q4 9.425 4 9V7q0-1.25.875-2.125T7 4h2q.425 0 .713.288T10 5q0 .425-.288.713T9 6H7q-.425 0-.713.288T6 7v2q0 .95-.55 1.725T4 11.825v.35q.9.325 1.45 1.1T6 15v2q0 .425.288.713T7 18h2q.425 0 .713.288T10 19q0 .425-.288.713T9 20H7Z"/>`),
		g.Group(children),
	)
}