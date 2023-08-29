package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 8.976l4.243-4.243l-1.415-1.414L12 6.147L9.172 3.32L7.757 4.733L12 8.976ZM5 12a1 1 0 0 1 1-1h12a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1Zm7 3.024l-4.243 4.243l1.415 1.414L12 17.853l2.828 2.828l1.415-1.414L12 15.024Z"/>`),
		g.Group(children),
	)
}