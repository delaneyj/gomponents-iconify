package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NineKSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 15H11V9H6.5v3.5h3v1h-3V15ZM8 11.5V10h1.5v1.5H8Zm5 3.5h1.5v-2.25L16.25 15H18l-2.25-3L18 9h-1.75l-1.75 2.25V9H13v6ZM3 21V3h18v18H3Z"/>`),
		g.Group(children),
	)
}