package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProductVirtual(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 3H2v6h1v11a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9h1V3zM4 5h16v2H4V5zm15 15H5V9h14v11zM7 16.45a1.82 1.82 0 0 1 1.82-1.82h.09a2.94 2.94 0 1 1 5.799-.92a2.27 2.27 0 0 1 1.08 4.29H7.87A1.811 1.811 0 0 1 7 16.45z"/>`),
		g.Group(children),
	)
}