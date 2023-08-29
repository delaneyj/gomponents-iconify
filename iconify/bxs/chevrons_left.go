package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m8.121 12l4.94-4.939l-2.122-2.122L3.879 12l7.06 7.061l2.122-2.122z"/><path fill="currentColor" d="M17.939 4.939L10.879 12l7.06 7.061l2.122-2.122L15.121 12l4.94-4.939z"/>`),
		g.Group(children),
	)
}