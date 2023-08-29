package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M13 5.07A7.002 7.002 0 0 1 12 19A7 7 0 0 1 7.262 6.847L5.847 5.432A9 9 0 1 0 11 3.055v6.03h2V5.072Z"/><path d="M7.707 8.707a1 1 0 0 0 0 1.414l2.829 2.829a1 1 0 0 0 1.414-1.414L9.12 8.707a1 1 0 0 0-1.414 0Z"/></g>`),
		g.Group(children),
	)
}