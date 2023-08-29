package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 6.879l-7.061 7.06l2.122 2.122L12 11.121l4.939 4.94l2.122-2.122z"/>`),
		g.Group(children),
	)
}