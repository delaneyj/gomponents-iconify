package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSortSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.854 2.146a.5.5 0 0 0-.708 0l-3 3a.5.5 0 1 0 .708.708L4 3.707V13.5a.5.5 0 0 0 1 0V3.707l2.146 2.147a.5.5 0 1 0 .708-.708l-3-3Zm6.298 11.714a.5.5 0 0 0 .696 0l3-2.9a.5.5 0 1 0-.696-.72L12 12.321v-9.82a.5.5 0 0 0-1 0v9.82l-2.152-2.08a.5.5 0 1 0-.696.718l3 2.9Z"/>`),
		g.Group(children),
	)
}