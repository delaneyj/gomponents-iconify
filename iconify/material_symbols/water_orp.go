package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterOrp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 22q-.425 0-.713-.288T6.5 21v-4q0-.425.288-.713T7.5 16H10q.425 0 .713.288T11 17v4q0 .425-.288.713T10 22H7.5Zm.5-1.5h1.5v-3H8v3Zm4 1.5v-6h3.5q.625 0 1.063.438T17 17.5v1q0 .45-.25.825t-.65.575L17 22h-1.5l-.85-2H13.5v2H12Zm6 0v-6h3.5q.625 0 1.063.438T23 17.5v1q0 .625-.438 1.063T21.5 20h-2v2H18Zm-4.5-3.5h2v-1h-2v1Zm6 0h2v-1h-2v1Zm-15 1.35q-1.175-1.125-1.838-2.675T2 13.8q0-2.5 1.988-5.438T10 2q4.025 3.425 6.013 6.363T18 13.8v.2H6.5q-.825 0-1.413.588T4.5 16v3.85Z"/>`),
		g.Group(children),
	)
}