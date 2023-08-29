package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChairOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20v-1q-1.25 0-2.125-.875T1 16v-5q0-1.25.875-2.125T4 8V6q0-1.25.875-2.125T7 3h10q1.25 0 2.125.875T20 6v2q1.25 0 2.125.875T23 11v5q0 1.25-.875 2.125T20 19v1q0 .425-.288.713T19 21q-.425 0-.713-.288T18 20v-1H6v1q0 .425-.288.713T5 21Zm-1-4h16q.425 0 .713-.288T21 16v-5q0-.425-.288-.713T20 10q-.425 0-.713.288T19 11v4H5v-4q0-.425-.288-.713T4 10q-.425 0-.713.288T3 11v5q0 .425.288.713T4 17Zm3-4h10v-2q0-.675.275-1.225T18 8.8V6q0-.425-.288-.713T17 5H7q-.425 0-.713.288T6 6v2.8q.45.425.725.975T7 11v2Zm5 0Zm0 4Zm0-2Z"/>`),
		g.Group(children),
	)
}