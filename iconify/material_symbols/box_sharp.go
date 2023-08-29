package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BoxSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 21V5.8L5.3 3h13.4L21 5.8V21H3ZM5.4 6h13.2l-.85-1H6.25L5.4 6ZM8 16l4-2l4 2V8H8v8Z"/>`),
		g.Group(children),
	)
}