package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpRightThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 3a1 1 0 1 0 0 2h10.585L3.297 27.288a1 1 0 1 0 1.414 1.415L27 6.414V17a1 1 0 1 0 2 0V4a1 1 0 0 0-1-1H15Z"/>`),
		g.Group(children),
	)
}