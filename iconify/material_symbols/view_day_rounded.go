package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewDayRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.975 20q-.425 0-.7-.288T3 19q0-.425.288-.713T4 18h16.025q.425 0 .7.288T21 19q0 .425-.288.713T20 20H3.975ZM5 16q-.825 0-1.413-.588T3 14v-4q0-.825.588-1.413T5 8h14q.825 0 1.413.588T21 10v4q0 .825-.588 1.413T19 16H5ZM3.975 6q-.425 0-.7-.288T3 5q0-.425.288-.713T4 4h16.025q.425 0 .7.288T21 5q0 .425-.288.713T20 6H3.975Z"/>`),
		g.Group(children),
	)
}