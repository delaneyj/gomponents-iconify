package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StickyNoteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 21H3V3h18v13l-5 5Zm-1-2l4-4h-4v4Zm-4-3h2v-6h3V8H8v2h3v6Z"/>`),
		g.Group(children),
	)
}