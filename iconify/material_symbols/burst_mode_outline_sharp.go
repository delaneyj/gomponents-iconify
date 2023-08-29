package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BurstModeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h2v14H1Zm4 0V5h2v14H5Zm4 0V5h14v14H9Zm2-2h10V7H11v10Zm1-2h8l-2.6-3.5l-1.9 2.5l-1.4-1.85L12 15Zm-1 2V7v10Z"/>`),
		g.Group(children),
	)
}