package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardOnscreenSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 17h8v-2H8v2Zm-3-3h2v-2H5v2Zm3 0h2v-2H8v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2ZM2 20V4h20v16H2ZM4 9h16V6H4v3Z"/>`),
		g.Group(children),
	)
}