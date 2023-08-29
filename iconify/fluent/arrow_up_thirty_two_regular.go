package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUpThirtyTwoRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 29a1 1 0 0 1-1-1V6.414l-8.293 8.293a1 1 0 0 1-1.414-1.414l10-10a1 1 0 0 1 1.414 0l10 10a1 1 0 0 1-1.414 1.414L17 6.414V28a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}