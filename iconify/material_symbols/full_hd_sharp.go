package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullHdSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.5 15H11v-2h1.5v2H14V9h-1.5v2.5H11V9H9.5v6Zm5.5 0h3.75l.75-.75v-4.5L18.75 9H15v6Zm1.5-1.5v-3H18v3h-1.5ZM4.5 15H6v-2h2v-1.5H6v-1h2.5V9h-4v6ZM1 20V4h22v16H1Z"/>`),
		g.Group(children),
	)
}