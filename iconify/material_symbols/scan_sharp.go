package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 9h5l-5-5v5ZM4 22v-5h16v5H4Zm-3-7v-2h22v2H1Zm3-4V2h10l6 6v3H4Z"/>`),
		g.Group(children),
	)
}