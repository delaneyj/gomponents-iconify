package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpAdUnits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 1H5v22h14V1zm-2 18H7V5h10v14z"/><path fill="currentColor" d="M8 6h8v2H8z"/>`),
		g.Group(children),
	)
}