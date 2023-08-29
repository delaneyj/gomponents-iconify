package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirlineSeatFlatRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 14q-.425 0-.713-.288T9 13V9q0-.825.588-1.413T11 7h7q1.65 0 2.825 1.175T22 11v2q0 .425-.288.713T21 14H10Zm11 3H3q-.425 0-.713-.288T2 16q0-.425.288-.713T3 15h18q.425 0 .713.288T22 16q0 .425-.288.713T21 17ZM5 14q-1.25 0-2.125-.875T2 11q0-1.25.875-2.125T5 8q1.25 0 2.125.875T8 11q0 1.25-.875 2.125T5 14Z"/>`),
		g.Group(children),
	)
}