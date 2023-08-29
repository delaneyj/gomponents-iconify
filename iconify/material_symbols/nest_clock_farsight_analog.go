package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestClockFarsightAnalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m14.55 16.55l1.4-1.425l-2.95-2.95V8h-2v5l3.55 3.55ZM11 6h2V4h-2v2Zm7 7h2v-2h-2v2Zm-7 7h2v-2h-2v2Zm-7-7h2v-2H4v2Zm8 9q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/>`),
		g.Group(children),
	)
}