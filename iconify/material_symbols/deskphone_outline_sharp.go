package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeskphoneOutlineSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 19h2V5h-2v14ZM6 14h2v-2H6v2Zm0 3h2v-2H6v2Zm0-6h8V7H6v4Zm3 3h2v-2H9v2Zm0 3h2v-2H9v2Zm3-3h2v-2h-2v2Zm0 3h2v-2h-2v2Zm3 1V6H5v12h10ZM3 20V4h12V3h6v18h-6v-1H3Zm2-2V6v12Z"/>`),
		g.Group(children),
	)
}