package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassEmptyRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 20h8v-3q0-1.65-1.175-2.825T12 13q-1.65 0-2.825 1.175T8 17v3Zm4-9q1.65 0 2.825-1.175T16 7V4H8v3q0 1.65 1.175 2.825T12 11ZM5 22q-.425 0-.713-.288T4 21q0-.425.288-.713T5 20h1v-3q0-1.525.713-2.863T8.7 12q-1.275-.8-1.987-2.138T6 7V4H5q-.425 0-.713-.288T4 3q0-.425.288-.713T5 2h14q.425 0 .713.288T20 3q0 .425-.288.713T19 4h-1v3q0 1.525-.713 2.863T15.3 12q1.275.8 1.988 2.138T18 17v3h1q.425 0 .713.288T20 21q0 .425-.288.713T19 22H5Z"/>`),
		g.Group(children),
	)
}