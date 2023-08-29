package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayShapes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m2 11l5-9l5 9H2Zm5 10q-1.65 0-2.825-1.175T3 17q0-1.65 1.175-2.825T7 13q1.65 0 2.825 1.175T11 17q0 1.65-1.175 2.825T7 21Zm6 0v-8h8v8h-8Zm1.4-10.5L13 9.1l2.6-2.6L13 3.9l1.4-1.4L17 5.1l2.6-2.6L21 3.9l-2.6 2.6L21 9.1l-1.4 1.4L17 7.9l-2.6 2.6Z"/>`),
		g.Group(children),
	)
}