package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwapVert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 13V5.825L5.425 8.4L4 7l5-5l5 5l-1.425 1.4L10 5.825V13H8Zm7 9l-5-5l1.425-1.4L14 18.175V11h2v7.175l2.575-2.575L20 17l-5 5Z"/>`),
		g.Group(children),
	)
}