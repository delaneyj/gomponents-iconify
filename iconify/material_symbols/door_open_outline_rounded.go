package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoorOpenOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 13q.425 0 .713-.288T12 12q0-.425-.288-.713T11 11q-.425 0-.713.288T10 12q0 .425.288.713T11 13Zm-4 8v-2l6-1V6.875q0-.375-.225-.675t-.575-.35L7 5V3l5.5.9q1.1.2 1.8 1.025T15 6.85v11.1q0 .725-.475 1.288t-1.2.687L7 21Zm0-2h10V5H7v14Zm-3 2q-.425 0-.713-.288T3 20q0-.425.288-.713T4 19h1V5q0-.825.588-1.413T7 3h10q.825 0 1.413.588T19 5v14h1q.425 0 .713.288T21 20q0 .425-.288.713T20 21H4Z"/>`),
		g.Group(children),
	)
}