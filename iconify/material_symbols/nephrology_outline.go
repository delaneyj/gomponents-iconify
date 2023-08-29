package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NephrologyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21v-4.175q-.25.1-.488.138T8 17q-2.5 0-4.25-1.75T2 11V9q0-2.5 1.75-4.25T8 3q1.25 0 2.125.875T11 6q0 1.25-.875 2.125T8 9H6V7h2q.425 0 .713-.288T9 6q0-.425-.288-.713T8 5Q6.35 5 5.175 6.175T4 9v2q0 1.65 1.175 2.825T8 15q.425 0 .713-.288T9 14q0-.425-.288-.713T8 13H6v-2h2q1.25 0 2.125.875T11 14v7H9Zm4 0v-7q0-1.25.875-2.125T16 11h2v2h-2q-.425 0-.713.288T15 14q0 .425.288.713T16 15q1.65 0 2.825-1.175T20 11V9q0-1.65-1.175-2.825T16 5q-.425 0-.713.288T15 6q0 .425.288.713T16 7h2v2h-2q-1.25 0-2.125-.875T13 6q0-1.25.875-2.125T16 3q2.5 0 4.25 1.75T22 9v2q0 2.5-1.75 4.25T16 17q-.275 0-.513-.038T15 16.825V21h-2ZM4 11V9v2Zm16-2v2v-2Z"/>`),
		g.Group(children),
	)
}