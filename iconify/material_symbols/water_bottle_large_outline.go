package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBottleLargeOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 20h10v-2h-2v-7h-2V9h4V7H7v2h2v7h2v2H7v2Zm0 2q-.825 0-1.413-.588T5 20v-2q0-.825.588-1.413T7 16v-5q-.825 0-1.413-.588T5 9V7q0-.825.588-1.413T7 5h3V4H9V2h6v2h-1v1h3q.825 0 1.413.588T19 7v2q0 .825-.588 1.413T17 11v5q.825 0 1.413.588T19 18v2q0 .825-.588 1.413T17 22H7Zm5-8.5Z"/>`),
		g.Group(children),
	)
}