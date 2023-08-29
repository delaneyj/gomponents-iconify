package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func InsertPageBreakSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22v-5h16v5H4Zm9-13h5l-5-5v5Zm-4 6v-2h6v2H9Zm8 0v-2h6v2h-6ZM1 15v-2h6v2H1Zm3-4V2h10l6 6v3H4Z"/>`),
		g.Group(children),
	)
}