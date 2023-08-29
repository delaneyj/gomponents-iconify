package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeyboardSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 19V5h20v14H2Zm6-3h8v-2H8v2Zm-3-3h2v-2H5v2Zm3 0h2v-2H8v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2Zm3 0h2v-2h-2v2ZM5 10h2V8H5v2Zm3 0h2V8H8v2Zm3 0h2V8h-2v2Zm3 0h2V8h-2v2Zm3 0h2V8h-2v2Z"/>`),
		g.Group(children),
	)
}