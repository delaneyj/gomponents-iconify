package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluidBalanceRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 20q-.825 0-1.413-.588T2 18V6q0-.825.588-1.413T4 4h16q.825 0 1.413.588T22 6h-7q-1.65 0-2.825 1.175T11 10v5q0 1.425.525 2.725T13.075 20H4Zm15 3q-.825 0-1.413-.588T17 21v-1.1q-1.725-.35-2.863-1.713T13 15v-5q0-.825.575-1.413T15 8h6q.825 0 1.413.588T23 10v5q0 1.825-1.137 3.188T19 19.9V21h2q.425 0 .713.288T22 22q0 .425-.288.713T21 23h-2Zm.75-9H21v-4h-6v2h.75q.825 0 1.563.375T18.55 13.4q.2.3.525.45t.675.15ZM10 11q.425 0 .713-.288T11 10q0-.425-.288-.713T10 9H7q-.425 0-.713.288T6 10q0 .425.288.713T7 11h3Zm0 4q.425 0 .713-.288T11 14q0-.425-.288-.713T10 13H7q-.425 0-.713.288T6 14q0 .425.288.713T7 15h3Z"/>`),
		g.Group(children),
	)
}