package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m16 12l2-5H8l-2 5l2 5h10l-2-5Zm-.954 3l-1.2-3l1.2-3H9.354l-1.2 3l1.2 3h5.692Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}