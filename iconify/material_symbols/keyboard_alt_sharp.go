package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardAltSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M1 21V4h22v17H1Zm7-4h8v-1H8v1Zm-3-3h2v-2H5v2Zm4 0h2v-2H9v2Zm4 0h2v-2h-2v2Zm4 0h2v-2h-2v2ZM5 10h2V8H5v2Zm4 0h2V8H9v2Zm4 0h2V8h-2v2Zm4 0h2V8h-2v2Z"/>`),
		g.Group(children),
	)
}