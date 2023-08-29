package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasketballHoop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21 2H3a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h2v-4h14v4h2a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2m-4 10h-2V9H9v3H7V9a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v3M7 16v3.5L8 23l2-2l2 2l2-2l2 2l1-3.5V16H7Z"/>`),
		g.Group(children),
	)
}