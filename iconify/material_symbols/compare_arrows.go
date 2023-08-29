package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompareArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8 20l-1.4-1.425L9.175 16H2v-2h7.175L6.6 11.425L8 10l5 5l-5 5Zm8-6l-5-5l5-5l1.4 1.425L14.825 8H22v2h-7.175l2.575 2.575L16 14Z"/>`),
		g.Group(children),
	)
}