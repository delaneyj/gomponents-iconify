package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrologySharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 23v-2H7.625l-.975-6H9v-1.175q-.25.1-.488.138T8 14q-2.5 0-4.25-1.75T2 8V7q0-2.5 1.75-4.25T8 1q1.25 0 2.125.875T11 4q0 1.25-.875 2.125T8 7H6V5h2q.425 0 .713-.288T9 4q0-.425-.288-.713T8 3Q6.35 3 5.175 4.175T4 7v1q0 1.65 1.175 2.825T8 12q.425 0 .713-.288T9 11q0-.425-.288-.713T8 10H6V8h2q1.25 0 2.125.875T11 11v4h2v-4q0-1.25.875-2.125T16 8h2v2h-2q-.425 0-.713.288T15 11q0 .425.288.713T16 12q1.65 0 2.825-1.175T20 8V7q0-1.65-1.175-2.825T16 3q-.425 0-.713.287T15 4q0 .425.288.713T16 5h2v2h-2q-1.25 0-2.125-.875T13 4q0-1.25.875-2.125T16 1q2.5 0 4.25 1.75T22 7v1q0 2.5-1.75 4.25T16 14q-.275 0-.513-.038T15 13.825V15h2.35l-.975 6H15v2h-2v-2h-2v2H9Z"/>`),
		g.Group(children),
	)
}