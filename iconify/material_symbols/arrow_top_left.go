package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowTopLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 20V10H7.825l3.6 3.6l-1.4 1.425L4 9l6-6l1.425 1.425L7.825 8H19v12h-2Z"/>`),
		g.Group(children),
	)
}