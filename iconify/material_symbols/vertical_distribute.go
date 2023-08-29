package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VerticalDistribute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 22v-2h20v2H2Zm5-8.5v-3h10v3H7ZM2 4V2h20v2H2Z"/>`),
		g.Group(children),
	)
}