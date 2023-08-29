package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineStartDiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 16.175L13.175 12L9 7.825L4.825 12L9 16.175ZM9 19l-7-7l7-7l6 6h7v2h-7l-6 6Zm0-7Z"/>`),
		g.Group(children),
	)
}