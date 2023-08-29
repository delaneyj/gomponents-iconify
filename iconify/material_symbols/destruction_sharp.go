package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DestructionSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 22v-8h18v8H3Zm3.525-9.5L1.5 9.65l6-1l-1.625-5.875L10.85 6.3l3-5.3l1 6l5.875-1.625L17.2 10.35L21 12.5H6.525Z"/>`),
		g.Group(children),
	)
}