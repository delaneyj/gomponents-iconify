package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackedEmailSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5 17V3h18v14H5Zm-4 4V6.5h2V19h16.5v2H1Zm13-8.725l7-4.85V5l-7 4.85L7 5v2.425l7 4.85Z"/>`),
		g.Group(children),
	)
}