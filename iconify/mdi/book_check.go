package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m16.75 22.16l-2.75-3L15.16 18l1.59 1.59L20.34 16l1.16 1.41l-4.75 4.75M6 22a2 2 0 0 1-2-2V4c0-1.11.89-2 2-2h1v7l2.5-1.5L12 9V2h6a2 2 0 0 1 2 2v9.34A6.005 6.005 0 0 0 12.8 22H6Z"/>`),
		g.Group(children),
	)
}