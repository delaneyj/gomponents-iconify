package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 13v-2h20v2H2Zm5-3V7h10v3H7Zm0 7v-3h10v3H7Z"/>`),
		g.Group(children),
	)
}