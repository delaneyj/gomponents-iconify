package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HexagonOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.8 21l-5.2-9l5.2-9h10.4l5.2 9l-5.2 9H6.8Zm1.15-2h8.1l4.025-7l-4.025-7h-8.1L3.9 12l4.05 7ZM12 12Z"/>`),
		g.Group(children),
	)
}