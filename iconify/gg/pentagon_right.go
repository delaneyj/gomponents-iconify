package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 12L6 7h10l2 5l-2 5H6l2-5Zm.954 3l1.2-3l-1.2-3h5.692l1.2 3l-1.2 3H8.954Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}