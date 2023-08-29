package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhpSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 15V9H11v2h2V9h1.5v6H13v-2.5h-2V15H9.5ZM3 15V9h5v4H4.5v2H3Zm1.5-3.5h2v-1h-2v1Zm12 3.5V9h5v4H18v2h-1.5Zm1.5-3.5h2v-1h-2v1Z"/>`),
		g.Group(children),
	)
}