package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextSelectMoveForwardWordOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 21v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm4 16v-2h2v2h-2Zm0-16V3h2v2h-2Zm-2 11l-1.4-1.4l1.575-1.6H11v-2h6.175L15.6 9.4L17 8l4 4l-4 4ZM5 19h2V5H5v14Zm-2 2V3h6v18H3Zm2-2h2h-2Z"/>`),
		g.Group(children),
	)
}