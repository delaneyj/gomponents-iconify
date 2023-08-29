package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElectricMopedSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m13 23l-6-3h4v-2l6 3h-4v2Zm-6-6q-1.25 0-2.125-.875T4 14H2V7h8v5h3.5L17 7.65V5h-3V3h5v5.35L14.5 14H10q0 1.25-.875 2.125T7 17Zm0-2q.425 0 .713-.288T8 14H6q0 .425.288.713T7 15ZM5 6V4h5v2H5Zm14 11q-1.25 0-2.125-.875T16 14q0-1.25.875-2.125T19 11q1.25 0 2.125.875T22 14q0 1.25-.875 2.125T19 17Zm0-2q.425 0 .713-.288T20 14q0-.425-.288-.713T19 13q-.425 0-.713.288T18 14q0 .425.288.713T19 15Z"/>`),
		g.Group(children),
	)
}