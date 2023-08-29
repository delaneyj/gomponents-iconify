package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SplitHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M3 5.016h18v1.968H3V5.016zm0 6V9h18v2.016H3zm0 7.968v-6h18v6H3z" fill="currentColor"/>`),
		g.Group(children),
	)
}