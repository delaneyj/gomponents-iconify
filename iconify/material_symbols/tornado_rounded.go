package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TornadoRounded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.9 8L2.75 6q-.575-1-.013-2t1.738-1h15.05q1.175 0 1.738 1t-.013 2L20.1 8H3.9Zm2.9 5l-1.75-3h13.9l-1.75 3H6.8Zm3.475 6L7.95 15h8.1l-2.325 4Q13.15 20 12 20t-1.725-1Z"/>`),
		g.Group(children),
	)
}