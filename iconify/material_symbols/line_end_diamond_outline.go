package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineEndDiamondOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 16.175L19.175 12L15 7.825L10.825 12L15 16.175ZM15 19l-6-6H2v-2h7l6-6l7 7l-7 7Zm0-7Z"/>`),
		g.Group(children),
	)
}