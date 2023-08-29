package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineWeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 20v-1h18v1H3Zm0-3v-2h18v2H3Zm0-4v-3h18v3H3Zm0-5V4h18v4H3Z"/>`),
		g.Group(children),
	)
}