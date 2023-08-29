package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OverviewKeyOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 19V5h14v14H1Zm2-2h10V7H3v10Zm14 2V5h2v14h-2Zm4 0V5h2v14h-2ZM3 17V7v10Z"/>`),
		g.Group(children),
	)
}