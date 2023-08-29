package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreencastSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20V4h20v16Zm12-4v-2h4v2Zm0-3v-2h4v2ZM4 11v7h16v-8h-6V8h6V6h-8v5Z"/>`),
		g.Group(children),
	)
}