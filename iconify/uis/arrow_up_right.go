package uis

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M17 6H7c-.6 0-1 .4-1 1s.4 1 1 1h7.6l-8.3 8.3c-.4.4-.4 1 0 1.4c.4.4 1 .4 1.4 0L16 9.4V17c0 .6.4 1 1 1s1-.4 1-1V7c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}