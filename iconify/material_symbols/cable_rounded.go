package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CableRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 21q-.425 0-.713-.288T4 20v-1h-.025q-.425 0-.7-.275t-.275-.7V15q0-.425.288-.713T4 14h1V7q0-1.65 1.175-2.825T9 3q1.65 0 2.825 1.175T13 7v10q0 .825.588 1.413T15 19q.825 0 1.413-.588T17 17v-7h-1q-.425 0-.713-.288T15 9V5.975q0-.425.275-.7t.7-.275H16V4q0-.425.288-.713T17 3h2q.425 0 .713.288T20 4v1h.025q.425 0 .7.275t.275.7V9q0 .425-.288.713T20 10h-1v7q0 1.65-1.175 2.825T15 21q-1.65 0-2.825-1.175T11 17V7q0-.825-.588-1.413T9 5q-.825 0-1.413.588T7 7v7h1q.425 0 .713.288T9 15v3.025q0 .425-.275.7t-.7.275H8v1q0 .425-.288.712T7 21H5Z"/>`),
		g.Group(children),
	)
}