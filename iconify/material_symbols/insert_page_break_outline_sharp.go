package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertPageBreakOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-5h2v3h12v-3h2v5H4Zm0-11V2h10l6 6v3h-2V9h-5V4H6v7H4Zm5 4v-2h6v2H9Zm8 0v-2h6v2h-6ZM1 15v-2h6v2H1Zm11-4Zm0 6Z"/>`),
		g.Group(children),
	)
}