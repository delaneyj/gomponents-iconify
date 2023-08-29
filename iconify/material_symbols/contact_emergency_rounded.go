package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContactEmergencyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 21V3h24v18H0Zm9-7q1.25 0 2.125-.875T12 11q0-1.25-.875-2.125T9 8q-1.25 0-2.125.875T6 11q0 1.25.875 2.125T9 14Zm-6.9 5h13.8q-1.05-1.875-2.9-2.938T9 15q-2.15 0-4 1.063T2.1 19Zm15.15-8.7v.95q0 .325.213.537T18 12q.325 0 .537-.213t.213-.537v-.95l.825.475q.275.15.575.075t.45-.35q.15-.275.075-.575t-.35-.45L19.5 9l.825-.475q.275-.15.35-.45T20.6 7.5q-.15-.275-.45-.35t-.575.075l-.825.475v-.95q0-.325-.213-.537T18 6q-.325 0-.537.213t-.213.537v.95l-.825-.475q-.275-.15-.575-.075t-.45.35q-.15.275-.075.575t.35.45L16.5 9l-.825.475q-.275.15-.35.45t.075.575q.15.275.45.35t.575-.075l.825-.475Z"/>`),
		g.Group(children),
	)
}