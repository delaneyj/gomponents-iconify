package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 16V7H8V5h11v11h-2Zm-5 5v-9H3v-2h11v11h-2Z"/>`),
		g.Group(children),
	)
}