package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeartPin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 20l-7.5-7.428A5 5 0 1 1 12 6.006a5 5 0 0 1 8.5 5.179m.621 8.936a3 3 0 1 0-4.242 0c.418.419 1.125 1.045 2.121 1.879c1.051-.89 1.759-1.516 2.121-1.879zM19 18v.01"/>`),
		g.Group(children),
	)
}