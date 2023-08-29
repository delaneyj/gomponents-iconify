package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoKSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15h1.5v-2.25L16.25 15H18l-2.25-3L18 9h-1.75l-1.75 2.25V9H13v6Zm-6.5 0H11v-1.5H8v-1h3V9H6.5v1.5h3v1h-3V15ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}