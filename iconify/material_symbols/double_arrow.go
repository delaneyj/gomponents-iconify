package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DoubleArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.05 19l5-7l-5-7H8.5l5 7l-5 7H6.05ZM12 19l5-7l-5-7h2.45l5 7l-5 7H12Z"/>`),
		g.Group(children),
	)
}