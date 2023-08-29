package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShelfPositionOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V3h19v18H3Zm2-5v3h15v-3H5Zm12-2h3V5h-3v9ZM5 14h3V5H5v9Zm5 0h5V5h-5v9Z"/>`),
		g.Group(children),
	)
}