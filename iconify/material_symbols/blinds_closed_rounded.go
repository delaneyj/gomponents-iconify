package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlindsClosedRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 19V5q0-.825.588-1.413T6 3h12q.825 0 1.413.588T20 5v14h1q.425 0 .713.288T22 20q0 .425-.288.713T21 21h-4.25q0 .725-.513 1.238T15 22.75q-.725 0-1.238-.513T13.25 21H3q-.425 0-.713-.288T2 20q0-.425.288-.713T3 19h1ZM6 7h8V5H6v2Zm10 0h2V5h-2v2ZM6 11h8V9H6v2Zm10 0h2V9h-2v2ZM6 15h8v-2H6v2Zm10 0h2v-2h-2v2ZM6 19h8v-2H6v2Zm10 0h2v-2h-2v2Z"/>`),
		g.Group(children),
	)
}