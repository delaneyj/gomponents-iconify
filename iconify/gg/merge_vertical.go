package gg

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MergeVertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8.976 12L4.733 7.757L3.32 9.172L6.147 12L3.32 14.828l1.414 1.415L8.976 12ZM12 19a1 1 0 0 1-1-1V6a1 1 0 1 1 2 0v12a1 1 0 0 1-1 1Zm3.024-7l4.243 4.243l1.414-1.415L17.853 12l2.828-2.828l-1.414-1.415L15.024 12Z"/>`),
		g.Group(children),
	)
}