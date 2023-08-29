package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteTwoSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 14h-5v5l5-5ZM3 21V3h18v12l-6 6H3Zm4-7h5v-2H7v2Zm0-4h10V8H7v2Z"/>`),
		g.Group(children),
	)
}