package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditDocumentSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 22V2h10l6 6v3.1l-8 7.975V22H4Zm10 0v-2.125l5.15-5.175l2.15 2.1l-5.175 5.2H14Zm8.025-5.9L19.9 13.975l1.425-1.425l2.075 2.175l-1.375 1.375ZM13 9h5l-5-5v5Z"/>`),
		g.Group(children),
	)
}