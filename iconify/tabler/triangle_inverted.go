package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleInverted(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.24 20.043L1.818 5.983A1.989 1.989 0 0 1 3.518 3h16.845a1.989 1.989 0 0 1 1.7 2.983l-8.422 14.06a1.989 1.989 0 0 1-3.4 0z"/>`),
		g.Group(children),
	)
}