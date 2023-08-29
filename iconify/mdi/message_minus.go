package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MessageMinus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2a2 2 0 0 1 2 2v12c0 1.11-.89 2-2 2H6l-4 4V4a2 2 0 0 1 2-2h16M8 9v2h8V9H8Z"/>`),
		g.Group(children),
	)
}