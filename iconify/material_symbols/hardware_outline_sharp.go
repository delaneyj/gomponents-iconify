package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HardwareOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 21V8H4q0-2.075 1.463-3.538T9 3h6v3l3-3h2v8h-2l-3-3v13H9Zm4-9Zm-2 7h2v-6h-2v6Zm0-8h2V5H9q-.65 0-1.225.263t-1 .737H11v5Zm2 0V5v6Zm0 8v-6v6Z"/>`),
		g.Group(children),
	)
}