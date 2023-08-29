package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 14V7h9q1.65 0 2.825 1.175T22 11v3H9Zm2-5v3v-3Zm-9 8v-2h20v2H2Zm3-3q-1.25 0-2.125-.875T2 11q0-1.25.875-2.125T5 8q1.25 0 2.125.875T8 11q0 1.25-.875 2.125T5 14Zm0-2q.425 0 .713-.288T6 11q0-.425-.288-.713T5 10q-.425 0-.713.288T4 11q0 .425.288.713T5 12Zm6 0h9v-1q0-.825-.588-1.413T18 9h-7v3Zm-6-1Z"/>`),
		g.Group(children),
	)
}