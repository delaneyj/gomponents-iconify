package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoOutlineRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.475 3h15.05q1.175 0 1.738 1t-.013 2l-7.525 13Q13.15 20 12 20t-1.725-1L2.75 6q-.575-1-.013-2t1.738-1Zm0 2L6.2 8h11.6l1.725-3H4.475Zm2.9 5L9.1 13h5.8l1.725-3h-9.25Zm2.9 5L12 18l1.725-3h-3.45Z"/>`),
		g.Group(children),
	)
}