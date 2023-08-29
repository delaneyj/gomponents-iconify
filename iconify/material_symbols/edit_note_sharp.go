package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EditNoteSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21v-2.125l5.3-5.3l2.125 2.125l-5.3 5.3H12Zm-9-5v-2h7v2H3Zm17.125-1L18 12.875l1.425-1.425l2.125 2.125L20.125 15ZM3 12v-2h11v2H3Zm0-4V6h11v2H3Z"/>`),
		g.Group(children),
	)
}