package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 20v-3h2V4h17v2H6v11h6v3H2Zm12 0V8h8v12h-8Zm2-3h4v-7h-4v7Zm0 0h4h-4Z"/>`),
		g.Group(children),
	)
}