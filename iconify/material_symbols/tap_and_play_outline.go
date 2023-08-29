package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapAndPlayOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 23v-2h1V6H7v6H5V3q0-.825.588-1.413T7 1h10q.825 0 1.413.588T19 3v18q0 .825-.588 1.413T17 23h-1ZM5 23v-2q.825 0 1.413.588T7 23H5Zm4 0q0-1.65-1.175-2.825T5 19v-2q2.5 0 4.25 1.75T11 23H9Zm4 0q0-3.35-2.325-5.675T5 15v-2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T15 23h-2ZM7 4h10V3H7v1Zm0 0V3v1Z"/>`),
		g.Group(children),
	)
}