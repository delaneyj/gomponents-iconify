package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LooksThreeOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 17h6V7H9v2h4v2h-2v2h2v2H9v2Zm-6 4V3h18v18H3Zm2-2h14V5H5v14Zm0 0V5v14Z"/>`),
		g.Group(children),
	)
}