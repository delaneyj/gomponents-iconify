package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.56 10.642L6.355 3.95l1.9 1.9a9.004 9.004 0 0 1 11.156 1.256l-1.414 1.415a7.003 7.003 0 0 0-8.28-1.21l1.537 1.538l-6.692 1.793Zm14.88 2.716l-1.794 6.692l-1.9-1.9A9.003 9.003 0 0 1 4.59 16.894l1.414-1.415a7.003 7.003 0 0 0 8.28 1.21l-1.537-1.538l6.692-1.793Z"/>`),
		g.Group(children),
	)
}