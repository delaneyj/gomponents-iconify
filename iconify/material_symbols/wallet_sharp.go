package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WalletSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2ZM4 8h16V6H4v2Zm11.775 6.075L20 10.525V10H4v1.225l11.775 2.85Z"/>`),
		g.Group(children),
	)
}