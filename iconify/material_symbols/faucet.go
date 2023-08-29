package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faucet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 15v-2h5v-1H5q-.425 0-.713-.288T4 11q0-.425.288-.713T5 10h2q.825 0 1.413.588T9 12v1h2V6.1q0-1.725 1.2-2.912T15.125 2q1.15 0 2.125.6t1.5 1.65l1.15 2.3l-1.8.9l-1.15-2.3q-.275-.525-.775-.838T15.1 4q-.875 0-1.487.613T13 6.1V13h2v-1q0-.825.588-1.413T17 10h2q.425 0 .713.288T20 11q0 .425-.288.713T19 12h-2v1h5v2H2Zm4 6q-.825 0-1.413-.588T4 19v-3h16v3q0 .825-.588 1.413T18 21H6Z"/>`),
		g.Group(children),
	)
}