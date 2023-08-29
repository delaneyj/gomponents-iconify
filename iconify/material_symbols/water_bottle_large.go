package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterBottleLarge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22q-.825 0-1.413-.588T5 20v-2q0-.825.588-1.413T7 16v-5q-.825 0-1.413-.588T5 9V7q0-.825.588-1.413T7 5h3V4H9V2h6v2h-1v1h3q.825 0 1.413.588T19 7v2q0 .825-.588 1.413T17 11v5q.825 0 1.413.588T19 18v2q0 .825-.588 1.413T17 22H7Z"/>`),
		g.Group(children),
	)
}