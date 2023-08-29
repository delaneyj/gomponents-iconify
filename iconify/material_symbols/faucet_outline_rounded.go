package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaucetOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 13v-1H5q-.425 0-.713-.288T4 11q0-.425.288-.713T5 10h2q.825 0 1.413.588T9 12v1h2V6.1q0-1.725 1.2-2.912T15.125 2q1.15 0 2.125.6t1.5 1.65l.7 1.4q.2.375.063.763T19 7q-.375.2-.763.063t-.587-.513l-.7-1.4q-.275-.525-.775-.838T15.1 4q-.875 0-1.487.613T13 6.1V13h2v-1q0-.825.588-1.413T17 10h2q.425 0 .713.288T20 11q0 .425-.288.713T19 12h-2v1h4q.425 0 .713.288T22 14q0 .425-.288.713T21 15H3q-.425 0-.713-.288T2 14q0-.425.288-.713T3 13h4Zm-1 8q-.825 0-1.413-.588T4 19v-4h2v4h12v-4h2v4q0 .825-.588 1.413T18 21H6Z"/>`),
		g.Group(children),
	)
}