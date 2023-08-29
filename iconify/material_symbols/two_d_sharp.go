package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoDSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M6.5 15H11v-1.5H8v-1h3V9H6.5v1.5h3v1h-3V15ZM3 21V3h18v18H3Zm10-6h4.25l.75-.75v-4.5L17.25 9H13v6Zm1.5-1.5v-3h2v3h-2Z"/>`),
		g.Group(children),
	)
}