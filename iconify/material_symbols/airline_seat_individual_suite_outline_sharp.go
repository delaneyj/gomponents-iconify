package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatIndividualSuiteOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 17V7h2v8h8V7h12v10H1Zm12-2h8V9h-8v6Zm0-6v6v-6Zm-6 5q1.25 0 2.125-.875T10 11q0-1.25-.875-2.125T7 8q-1.25 0-2.125.875T4 11q0 1.25.875 2.125T7 14Zm0-2q-.425 0-.713-.288T6 11q0-.425.288-.713T7 10q.425 0 .713.288T8 11q0 .425-.288.713T7 12Zm0-1Z"/>`),
		g.Group(children),
	)
}