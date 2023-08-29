package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterAltOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 20v-7L2.95 4h18.1L14 13v7h-4Zm2-7.7L16.95 6h-9.9L12 12.3Zm0 0Z"/>`),
		g.Group(children),
	)
}