package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Notes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 18v-2h12v2H3Zm0-5v-2h18v2H3Zm0-5V6h18v2H3Z"/>`),
		g.Group(children),
	)
}