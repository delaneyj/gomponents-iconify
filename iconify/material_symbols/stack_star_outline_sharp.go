package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackStarOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.725 18.5L15 17.125l2.275 1.375l-.6-2.6l2-1.725l-2.625-.225L15 11.5l-1.05 2.45l-2.625.225l2 1.725l-.6 2.6ZM6 14v2H2V2h14v4h-2V4H4v10h2Zm2 8V8h14v14H8Zm2-2h10V10H10v10Zm5-5Z"/>`),
		g.Group(children),
	)
}