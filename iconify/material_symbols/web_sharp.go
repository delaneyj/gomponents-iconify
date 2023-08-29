package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16H2Zm2-2h10.5v-3.5H4V18Zm12.5 0H20V9h-3.5v9ZM4 12.5h10.5V9H4v3.5Z"/>`),
		g.Group(children),
	)
}