package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDobuleDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.613 13.935L0 2.344L2.352 0l9.261 9.239L20.874 0l2.352 2.344zm0 10.065L0 12.409l2.352-2.344l9.261 9.239l9.261-9.239l2.352 2.344z"/>`),
		g.Group(children),
	)
}