package uim

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.608 8.01a1.004 1.004 0 0 0-.216-1.089l-4.628-4.628a1 1 0 0 0-1.414 1.414l2.921 2.921H8.314a4.005 4.005 0 0 0-4 4V21a1 1 0 0 0 2 0V10.628a2.002 2.002 0 0 1 2-2h7.958L13.35 11.55a1 1 0 1 0 1.414 1.414l4.628-4.628a1 1 0 0 0 .216-.325Z"/>`),
		g.Group(children),
	)
}