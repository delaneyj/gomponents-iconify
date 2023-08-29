package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FortRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19v-1.175q0-.4.15-.763t.425-.637L3 15V9L1.575 7.575Q1.3 7.3 1.15 6.937T1 6.175V4q0-.425.288-.713T2 3q.425 0 .713.288T3 4v1h2V4q0-.425.288-.713T6 3q.425 0 .713.288T7 4v1h2V4q0-.425.288-.713T10 3q.425 0 .713.288T11 4v2.175q0 .4-.15.763t-.425.637L9 9v1h6V9l-1.425-1.425q-.275-.275-.425-.638T13 6.175V4q0-.425.287-.713T14 3q.425 0 .713.288T15 4v1h2V4q0-.425.288-.713T18 3q.425 0 .713.288T19 4v1h2V4q0-.425.288-.713T22 3q.425 0 .713.288T23 4v2.175q0 .4-.15.763t-.425.637L21 9v6l1.425 1.425q.275.275.425.638t.15.762V19q0 .825-.588 1.413T21 21h-6q-.425 0-.713-.288T14 20v-2q0-.825-.588-1.413T12 16q-.825 0-1.413.588T10 18v2q0 .425-.288.713T9 21H3q-.825 0-1.413-.588T1 19Z"/>`),
		g.Group(children),
	)
}